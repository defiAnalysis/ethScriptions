package util

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"math/rand"
	"models"
	"runtime"
	"strings"
	"time"
	"utils"
	"utils/redis"
)

func init() {
	beego.LoadAppConfig("ini", "../conf/app.conf")
}

type BlockAnalysis struct {
	block       chan uint64
	transaction chan TransactionStr

	chainid     string
	lastheight  uint64
	blockheight uint64
	client      *ethclient.Client

	inputData chan InputData

	validTransactions chan InputData

	session orm.Ormer
}

type InputData struct {
	DecodeData string          `json:"decode_data"`
	Creator    common.Address  `json:"creator"`
	Owner      *common.Address `json:"owner"`
	Hash       common.Hash     `json:"hash"`
	BlockTime  uint64          `json:"block_time"`
	Content    string          `json:"content"`
}

type TransactionStr struct {
	transaction types.Transaction
	block       types.Block
}

func (this *BlockAnalysis) Run() {
	this.chainid = "ETH"

	this.lastheight, _ = redis.RedisGet("height:" + this.chainid).Uint64()

	client, err := ethclient.Dial(beego.AppConfig.String("rpc::ETH"))

	if err != nil {
		beego.Error("节点客户端初始化异常：", err)
	} else {
		this.client = client

		blockheight, err := this.client.BlockNumber(context.Background())
		if err != nil {
			beego.Error("节点客户端初始化异常：", err)
		}
		this.blockheight = blockheight
	}

	if this.lastheight == 0 {
		this.lastheight = this.blockheight - 1
		redis.RedisSet("height:"+this.chainid, this.lastheight, -1)
	}

	this.transaction = make(chan TransactionStr, 5000)
	this.block = make(chan uint64, 2000)

	this.inputData = make(chan InputData, 4000)
	this.validTransactions = make(chan InputData, 8000)

	this.session = orm.NewOrm()

	go this.monitorHeight()
	go this.analysisBlock()
	go this.analysisTransaction()
	go this.validTransaction()
	//定时初始化以太坊客户端

	go this.DecodeData()

	go func() {
		for {
			this.initClient()
			time.Sleep(30 * time.Second)
		}
	}()
	select {}
}

func (this *BlockAnalysis) initClient() {
	rpcs := strings.Split(beego.AppConfig.String("rpcs::ETH"), ",")

	for {
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(rpcs))
		client, err := ethclient.Dial(rpcs[index])
		if err != nil {
			log.Println("节点客户端初始化异常：", err)
			time.Sleep(3 * time.Second)

			continue
		} else {
			_, err := client.BlockNumber(context.Background())
			if err != nil {
				time.Sleep(3 * time.Second)

				continue
			}

			this.client = client
			break
		}
	}
}

// 扫描区块
func (this *BlockAnalysis) monitorHeight() {
	for {
		beego.Info("this.lastheight：", this.lastheight)
		beego.Info("this.blockheight：", this.blockheight)
		if this.lastheight < this.blockheight {
			if this.lastheight+100 >= this.blockheight {
				start := this.lastheight
				end := this.lastheight + (this.blockheight - this.lastheight)
				this.monitorBlock(start, end)
				this.lastheight = end + 1
			} else {
				start := this.lastheight
				end := this.lastheight + 100
				this.monitorBlock(start, end)
				this.lastheight += 100
			}

			go redis.RedisSet("height:"+this.chainid, this.lastheight, -1)
		}

		this.blockheight, _ = this.client.BlockNumber(context.Background())
		time.Sleep(500 * time.Millisecond)
	}
}

func (this *BlockAnalysis) monitorBlock(start, end uint64) {
	for start <= end {
		time.Sleep(200 * time.Millisecond)
		this.block <- start
		start++
	}
}

// 解析区块
func (this *BlockAnalysis) analysisBlock() {
	rpcs := strings.Split(beego.AppConfig.String("rpcs::ETH"), ",")
	for {
		if runtime.NumGoroutine() < 1000 {
			blockheight := <-this.block
			beego.Info("[", this.chainid, "]", "[blockheight]", blockheight)

			go func() {
			A:
				for {
					rand.Seed(time.Now().UnixNano())
					index := rand.Intn(len(rpcs))
					client, err := ethclient.Dial(rpcs[index])
					if err != nil {
						log.Println("节点客户端初始化异常：", err)
						time.Sleep(200 * time.Millisecond)

						continue A
					} else {
						if block, e := client.BlockByNumber(context.Background(), big.NewInt(int64(blockheight))); e != nil {
							log.Println("区块解析失败，打回：", blockheight)
							beego.Error("e:", e.Error())
							time.Sleep(200 * time.Millisecond)
							continue A
						} else {
							if block != nil {
								log.Println("Height: ", "[", blockheight, "]")
								for _, tran := range block.Transactions() {
									if tran == nil {
										continue
									}

									tranData := TransactionStr{
										transaction: *tran,
										block:       *block,
									}

									this.transaction <- tranData
								}

								break A
							} else {
								continue A
							}
						}
					}
				}
			}()
		}
	}
}

// 解析交易
func (this *BlockAnalysis) analysisTransaction() {

	for {
		tx := <-this.transaction
		if len(tx.transaction.Data()) == 0 {
			continue
		}

		go func(tranStr TransactionStr) {
			tran := tranStr.transaction

			inputData := hex.EncodeToString(tran.Data())

			str, err := utils.HexToString(inputData)
			if err != nil {
				fmt.Println("解析失败:", err)
				return
			}

			flag2 := false
			if strings.HasPrefix(str, "data:,") {
				fmt.Println("字符串以'data:,'开头")
				flag2 = true
			}

			if flag2 {
				from, err := GetFrom(tran)
				if err != nil {
					beego.Info("GetFrom err：", err.Error())
				}

				data := InputData{
					DecodeData: str,
					Creator:    common.HexToAddress(from),
					Owner:      tran.To(),
					Hash:       tran.Hash(),
					BlockTime:  tranStr.block.Time(),
					Content:    strings.TrimPrefix(str, "data:,"),
				}

				beego.Info("data:", data)

				this.validTransactions <- data
			}

		}(tx)
	}
}

func findAddr(adds []string, addr string) bool {
	flag := false
	for i := 0; i < len(adds); i++ {
		if adds[i] == addr {
			flag = true
			break
		}
	}

	return flag
}

// 校验交易
func (this *BlockAnalysis) validTransaction() {
	rpcs := strings.Split(beego.AppConfig.String("rpcs::ETH"), ",")
	for {
		data := <-this.validTransactions

		beego.Info("data:", data)

		go func(data InputData) {
			for {
				rand.Seed(time.Now().UnixNano())
				index := rand.Intn(len(rpcs))
				client, err := ethclient.Dial(rpcs[index])
				if err != nil {
					log.Println("节点客户端初始化异常：", err)
					time.Sleep(200 * time.Millisecond)

					continue
				} else {
					if receipt, e := client.TransactionReceipt(context.Background(), data.Hash); e != nil {
						time.Sleep(200 * time.Millisecond)
						log.Println("e:", e)
						continue
					} else {
						if receipt.Status == 1 {
							beego.Info("发送data:", data)
							this.inputData <- data
							return
						} else {
							return
						}
					}
				}

				return
			}
		}(data)

	}

}

// 打怪开始
func (this *BlockAnalysis) DecodeData() {
	for {
		inputData := <-this.inputData
		beego.Info("inputData:", inputData)

		ethScription := models.EthScription{
			DecodeData: inputData.DecodeData,
		}

		if err := this.session.Read(&ethScription, "decode_data"); err != nil && err != orm.ErrNoRows {
			beego.Error("session Read err:", err.Error())

			continue
		} else if err != nil && err == orm.ErrNoRows {
			ethScripts := models.EthScription{
				DecodeData: inputData.DecodeData,
				Creator:    inputData.Creator.String(),
				Owner:      inputData.Owner.String(),
				Hash:       inputData.Hash.String(),
				BlockTime:  inputData.BlockTime,
				Content:    inputData.Content,
			}

			if _, err := this.session.Insert(&ethScripts); err != nil {
				beego.Error("err:", err)
			}
		}
	}
}

func GetFrom(tx types.Transaction) (string, error) {
	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), &tx)
	return from.String(), err
}
