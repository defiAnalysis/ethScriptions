// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package battle

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BattleBattleStruct is an auto generated low-level Go binding around an user-defined struct.
type BattleBattleStruct struct {
	Id      *big.Int
	TokenId *big.Int
	Exp     *big.Int
	Reward  *big.Int
}

// BattleFightHistory is an auto generated low-level Go binding around an user-defined struct.
type BattleFightHistory struct {
	TokeneId    *big.Int
	Monster     uint8
	Result      uint8
	Reward      *big.Int
	Exp         *big.Int
	BlockNUmber *big.Int
	Time        *big.Int
}

// BattleMonsterInfo is an auto generated low-level Go binding around an user-defined struct.
type BattleMonsterInfo struct {
	Monster     uint8
	Name        string
	WinRate     *big.Int
	RewardRange *big.Int
	ExpStart    *big.Int
	ExpEnd      *big.Int
}

// BattleUserChooseZoan is an auto generated low-level Go binding around an user-defined struct.
type BattleUserChooseZoan struct {
	Zoon     CryptoZooNFTZoonInfo
	Physical *big.Int
	Total    *big.Int
	Skill    *big.Int
}

// CryptoZooNFTCryptoZoon is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTCryptoZoon struct {
	Tribe      uint8
	Generation *big.Int
	Exp        *big.Int
	Dna        *big.Int
	FarmTime   *big.Int
	BornTime   *big.Int
}

// CryptoZooNFTZoonInfo is an auto generated low-level Go binding around an user-defined struct.
type CryptoZooNFTZoonInfo struct {
	Zoon    CryptoZooNFTCryptoZoon
	Tokenid *big.Int
}

// BattleMetaData contains all meta data concerning the Battle contract.
var BattleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumBattle.MonsterLevel\",\"name\":\"monster\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumBattle.BattleResult\",\"name\":\"result\",\"type\":\"uint8\"}],\"name\":\"BattleEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"}],\"name\":\"Evolve\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"_tribe\",\"type\":\"uint8\"}],\"name\":\"Spawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"Id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockTime\",\"type\":\"uint256\"}],\"name\":\"UserBattleEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amound\",\"type\":\"uint256\"}],\"name\":\"WithdrawEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"FightHistorys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokeneId\",\"type\":\"uint256\"},{\"internalType\":\"enumBattle.MonsterLevel\",\"name\":\"monster\",\"type\":\"uint8\"},{\"internalType\":\"enumBattle.BattleResult\",\"name\":\"result\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNUmber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"UserBattleTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"id\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"internalType\":\"structBattle.BattleStruct\",\"name\":\"_battleinfo\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"name\":\"Userbattle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumBattle.MonsterLevel\",\"name\":\"_monster\",\"type\":\"uint8\"}],\"name\":\"battle1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blocklists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chooseMonster\",\"outputs\":[{\"components\":[{\"internalType\":\"enumBattle.MonsterLevel\",\"name\":\"monster\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"winRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRange\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expEnd\",\"type\":\"uint256\"}],\"internalType\":\"structBattle.MonsterInfo[]\",\"name\":\"monster\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"chooseZoan\",\"outputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"enumCryptoZooNFT.Tribe\",\"name\":\"tribe\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"generation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dna\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"farmTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bornTime\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.CryptoZoon\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"tokenid\",\"type\":\"uint256\"}],\"internalType\":\"structCryptoZooNFT.ZoonInfo\",\"name\":\"zoon\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"physical\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"skill\",\"type\":\"uint256\"}],\"internalType\":\"structBattle.UserChooseZoan[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"day\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dayBattle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dogCoin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"evolveEgg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"generateDNA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getBattleTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getDayBattleTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_times\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEvolveEggPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getFightHistorys\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokeneId\",\"type\":\"uint256\"},{\"internalType\":\"enumBattle.MonsterLevel\",\"name\":\"monster\",\"type\":\"uint8\"},{\"internalType\":\"enumBattle.BattleResult\",\"name\":\"result\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNUmber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"internalType\":\"structBattle.FightHistory[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getPendingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getSpendNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"}],\"name\":\"getUserNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_userNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getZbattle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getxBattle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"layRnd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"layRndDogCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"layRndUsdt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractManagerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"randomExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_zoonerERC20\",\"type\":\"address\"},{\"internalType\":\"contractCryptoZooNFT\",\"name\":\"zoonerNFT\",\"type\":\"address\"}],\"name\":\"setBattleInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_user\",\"type\":\"address[]\"}],\"name\":\"setBlack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dogCoin\",\"type\":\"address\"}],\"name\":\"setTokenAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenids\",\"type\":\"uint256[]\"}],\"name\":\"setTokenids\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sign\",\"type\":\"address\"}],\"name\":\"signAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"suplusTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lasttime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_exp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"rs\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ss\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"vs\",\"type\":\"uint8\"}],\"name\":\"validSignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withDrawRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoonerERC20\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zoonerNFT\",\"outputs\":[{\"internalType\":\"contractCryptoZooNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BattleABI is the input ABI used to generate the binding from.
// Deprecated: Use BattleMetaData.ABI instead.
var BattleABI = BattleMetaData.ABI

// Battle is an auto generated Go binding around an Ethereum contract.
type Battle struct {
	BattleCaller     // Read-only binding to the contract
	BattleTransactor // Write-only binding to the contract
	BattleFilterer   // Log filterer for contract events
}

// BattleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BattleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BattleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BattleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BattleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BattleSession struct {
	Contract     *Battle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BattleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BattleCallerSession struct {
	Contract *BattleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BattleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BattleTransactorSession struct {
	Contract     *BattleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BattleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BattleRaw struct {
	Contract *Battle // Generic contract binding to access the raw methods on
}

// BattleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BattleCallerRaw struct {
	Contract *BattleCaller // Generic read-only contract binding to access the raw methods on
}

// BattleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BattleTransactorRaw struct {
	Contract *BattleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBattle creates a new instance of Battle, bound to a specific deployed contract.
func NewBattle(address common.Address, backend bind.ContractBackend) (*Battle, error) {
	contract, err := bindBattle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Battle{BattleCaller: BattleCaller{contract: contract}, BattleTransactor: BattleTransactor{contract: contract}, BattleFilterer: BattleFilterer{contract: contract}}, nil
}

// NewBattleCaller creates a new read-only instance of Battle, bound to a specific deployed contract.
func NewBattleCaller(address common.Address, caller bind.ContractCaller) (*BattleCaller, error) {
	contract, err := bindBattle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BattleCaller{contract: contract}, nil
}

// NewBattleTransactor creates a new write-only instance of Battle, bound to a specific deployed contract.
func NewBattleTransactor(address common.Address, transactor bind.ContractTransactor) (*BattleTransactor, error) {
	contract, err := bindBattle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BattleTransactor{contract: contract}, nil
}

// NewBattleFilterer creates a new log filterer instance of Battle, bound to a specific deployed contract.
func NewBattleFilterer(address common.Address, filterer bind.ContractFilterer) (*BattleFilterer, error) {
	contract, err := bindBattle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BattleFilterer{contract: contract}, nil
}

// bindBattle binds a generic wrapper to an already deployed contract.
func bindBattle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BattleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Battle *BattleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Battle.Contract.BattleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Battle *BattleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battle.Contract.BattleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Battle *BattleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Battle.Contract.BattleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Battle *BattleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Battle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Battle *BattleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Battle *BattleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Battle.Contract.contract.Transact(opts, method, params...)
}

// FightHistorys is a free data retrieval call binding the contract method 0x93959c21.
//
// Solidity: function FightHistorys(address , uint256 ) view returns(uint256 tokeneId, uint8 monster, uint8 result, uint256 reward, uint256 exp, uint256 blockNUmber, uint256 time)
func (_Battle *BattleCaller) FightHistorys(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TokeneId    *big.Int
	Monster     uint8
	Result      uint8
	Reward      *big.Int
	Exp         *big.Int
	BlockNUmber *big.Int
	Time        *big.Int
}, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "FightHistorys", arg0, arg1)

	outstruct := new(struct {
		TokeneId    *big.Int
		Monster     uint8
		Result      uint8
		Reward      *big.Int
		Exp         *big.Int
		BlockNUmber *big.Int
		Time        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokeneId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Monster = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Result = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Reward = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Exp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.BlockNUmber = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// FightHistorys is a free data retrieval call binding the contract method 0x93959c21.
//
// Solidity: function FightHistorys(address , uint256 ) view returns(uint256 tokeneId, uint8 monster, uint8 result, uint256 reward, uint256 exp, uint256 blockNUmber, uint256 time)
func (_Battle *BattleSession) FightHistorys(arg0 common.Address, arg1 *big.Int) (struct {
	TokeneId    *big.Int
	Monster     uint8
	Result      uint8
	Reward      *big.Int
	Exp         *big.Int
	BlockNUmber *big.Int
	Time        *big.Int
}, error) {
	return _Battle.Contract.FightHistorys(&_Battle.CallOpts, arg0, arg1)
}

// FightHistorys is a free data retrieval call binding the contract method 0x93959c21.
//
// Solidity: function FightHistorys(address , uint256 ) view returns(uint256 tokeneId, uint8 monster, uint8 result, uint256 reward, uint256 exp, uint256 blockNUmber, uint256 time)
func (_Battle *BattleCallerSession) FightHistorys(arg0 common.Address, arg1 *big.Int) (struct {
	TokeneId    *big.Int
	Monster     uint8
	Result      uint8
	Reward      *big.Int
	Exp         *big.Int
	BlockNUmber *big.Int
	Time        *big.Int
}, error) {
	return _Battle.Contract.FightHistorys(&_Battle.CallOpts, arg0, arg1)
}

// UserBattleTimes is a free data retrieval call binding the contract method 0x3fd6725d.
//
// Solidity: function UserBattleTimes(uint256 _tokenid) view returns(uint256)
func (_Battle *BattleCaller) UserBattleTimes(opts *bind.CallOpts, _tokenid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "UserBattleTimes", _tokenid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBattleTimes is a free data retrieval call binding the contract method 0x3fd6725d.
//
// Solidity: function UserBattleTimes(uint256 _tokenid) view returns(uint256)
func (_Battle *BattleSession) UserBattleTimes(_tokenid *big.Int) (*big.Int, error) {
	return _Battle.Contract.UserBattleTimes(&_Battle.CallOpts, _tokenid)
}

// UserBattleTimes is a free data retrieval call binding the contract method 0x3fd6725d.
//
// Solidity: function UserBattleTimes(uint256 _tokenid) view returns(uint256)
func (_Battle *BattleCallerSession) UserBattleTimes(_tokenid *big.Int) (*big.Int, error) {
	return _Battle.Contract.UserBattleTimes(&_Battle.CallOpts, _tokenid)
}

// Blocklists is a free data retrieval call binding the contract method 0x95361b13.
//
// Solidity: function blocklists(address ) view returns(bool)
func (_Battle *BattleCaller) Blocklists(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "blocklists", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Blocklists is a free data retrieval call binding the contract method 0x95361b13.
//
// Solidity: function blocklists(address ) view returns(bool)
func (_Battle *BattleSession) Blocklists(arg0 common.Address) (bool, error) {
	return _Battle.Contract.Blocklists(&_Battle.CallOpts, arg0)
}

// Blocklists is a free data retrieval call binding the contract method 0x95361b13.
//
// Solidity: function blocklists(address ) view returns(bool)
func (_Battle *BattleCallerSession) Blocklists(arg0 common.Address) (bool, error) {
	return _Battle.Contract.Blocklists(&_Battle.CallOpts, arg0)
}

// ChooseMonster is a free data retrieval call binding the contract method 0x64046e91.
//
// Solidity: function chooseMonster() view returns((uint8,string,uint256,uint256,uint256,uint256)[] monster, uint256 cost)
func (_Battle *BattleCaller) ChooseMonster(opts *bind.CallOpts) (struct {
	Monster []BattleMonsterInfo
	Cost    *big.Int
}, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "chooseMonster")

	outstruct := new(struct {
		Monster []BattleMonsterInfo
		Cost    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Monster = *abi.ConvertType(out[0], new([]BattleMonsterInfo)).(*[]BattleMonsterInfo)
	outstruct.Cost = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ChooseMonster is a free data retrieval call binding the contract method 0x64046e91.
//
// Solidity: function chooseMonster() view returns((uint8,string,uint256,uint256,uint256,uint256)[] monster, uint256 cost)
func (_Battle *BattleSession) ChooseMonster() (struct {
	Monster []BattleMonsterInfo
	Cost    *big.Int
}, error) {
	return _Battle.Contract.ChooseMonster(&_Battle.CallOpts)
}

// ChooseMonster is a free data retrieval call binding the contract method 0x64046e91.
//
// Solidity: function chooseMonster() view returns((uint8,string,uint256,uint256,uint256,uint256)[] monster, uint256 cost)
func (_Battle *BattleCallerSession) ChooseMonster() (struct {
	Monster []BattleMonsterInfo
	Cost    *big.Int
}, error) {
	return _Battle.Contract.ChooseMonster(&_Battle.CallOpts)
}

// ChooseZoan is a free data retrieval call binding the contract method 0x99e7ffaa.
//
// Solidity: function chooseZoan(address _addr) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256)[] result)
func (_Battle *BattleCaller) ChooseZoan(opts *bind.CallOpts, _addr common.Address) ([]BattleUserChooseZoan, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "chooseZoan", _addr)

	if err != nil {
		return *new([]BattleUserChooseZoan), err
	}

	out0 := *abi.ConvertType(out[0], new([]BattleUserChooseZoan)).(*[]BattleUserChooseZoan)

	return out0, err

}

// ChooseZoan is a free data retrieval call binding the contract method 0x99e7ffaa.
//
// Solidity: function chooseZoan(address _addr) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256)[] result)
func (_Battle *BattleSession) ChooseZoan(_addr common.Address) ([]BattleUserChooseZoan, error) {
	return _Battle.Contract.ChooseZoan(&_Battle.CallOpts, _addr)
}

// ChooseZoan is a free data retrieval call binding the contract method 0x99e7ffaa.
//
// Solidity: function chooseZoan(address _addr) view returns((((uint8,uint256,uint256,uint256,uint256,uint256),uint256),uint256,uint256,uint256)[] result)
func (_Battle *BattleCallerSession) ChooseZoan(_addr common.Address) ([]BattleUserChooseZoan, error) {
	return _Battle.Contract.ChooseZoan(&_Battle.CallOpts, _addr)
}

// Day is a free data retrieval call binding the contract method 0x7b76ac91.
//
// Solidity: function day() view returns(uint256)
func (_Battle *BattleCaller) Day(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "day")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Day is a free data retrieval call binding the contract method 0x7b76ac91.
//
// Solidity: function day() view returns(uint256)
func (_Battle *BattleSession) Day() (*big.Int, error) {
	return _Battle.Contract.Day(&_Battle.CallOpts)
}

// Day is a free data retrieval call binding the contract method 0x7b76ac91.
//
// Solidity: function day() view returns(uint256)
func (_Battle *BattleCallerSession) Day() (*big.Int, error) {
	return _Battle.Contract.Day(&_Battle.CallOpts)
}

// DayBattle is a free data retrieval call binding the contract method 0x34961bb1.
//
// Solidity: function dayBattle(uint256 , uint256 ) view returns(uint256)
func (_Battle *BattleCaller) DayBattle(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "dayBattle", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DayBattle is a free data retrieval call binding the contract method 0x34961bb1.
//
// Solidity: function dayBattle(uint256 , uint256 ) view returns(uint256)
func (_Battle *BattleSession) DayBattle(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Battle.Contract.DayBattle(&_Battle.CallOpts, arg0, arg1)
}

// DayBattle is a free data retrieval call binding the contract method 0x34961bb1.
//
// Solidity: function dayBattle(uint256 , uint256 ) view returns(uint256)
func (_Battle *BattleCallerSession) DayBattle(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Battle.Contract.DayBattle(&_Battle.CallOpts, arg0, arg1)
}

// DogCoin is a free data retrieval call binding the contract method 0x3827f276.
//
// Solidity: function dogCoin() view returns(address)
func (_Battle *BattleCaller) DogCoin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "dogCoin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DogCoin is a free data retrieval call binding the contract method 0x3827f276.
//
// Solidity: function dogCoin() view returns(address)
func (_Battle *BattleSession) DogCoin() (common.Address, error) {
	return _Battle.Contract.DogCoin(&_Battle.CallOpts)
}

// DogCoin is a free data retrieval call binding the contract method 0x3827f276.
//
// Solidity: function dogCoin() view returns(address)
func (_Battle *BattleCallerSession) DogCoin() (common.Address, error) {
	return _Battle.Contract.DogCoin(&_Battle.CallOpts)
}

// GenerateDNA is a free data retrieval call binding the contract method 0x47a849b6.
//
// Solidity: function generateDNA(uint256 _tokenId) view returns(uint256)
func (_Battle *BattleCaller) GenerateDNA(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "generateDNA", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateDNA is a free data retrieval call binding the contract method 0x47a849b6.
//
// Solidity: function generateDNA(uint256 _tokenId) view returns(uint256)
func (_Battle *BattleSession) GenerateDNA(_tokenId *big.Int) (*big.Int, error) {
	return _Battle.Contract.GenerateDNA(&_Battle.CallOpts, _tokenId)
}

// GenerateDNA is a free data retrieval call binding the contract method 0x47a849b6.
//
// Solidity: function generateDNA(uint256 _tokenId) view returns(uint256)
func (_Battle *BattleCallerSession) GenerateDNA(_tokenId *big.Int) (*big.Int, error) {
	return _Battle.Contract.GenerateDNA(&_Battle.CallOpts, _tokenId)
}

// GetBattleTimes is a free data retrieval call binding the contract method 0xf6ac6c71.
//
// Solidity: function getBattleTimes(uint256 _tokenId) view returns(uint256)
func (_Battle *BattleCaller) GetBattleTimes(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getBattleTimes", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBattleTimes is a free data retrieval call binding the contract method 0xf6ac6c71.
//
// Solidity: function getBattleTimes(uint256 _tokenId) view returns(uint256)
func (_Battle *BattleSession) GetBattleTimes(_tokenId *big.Int) (*big.Int, error) {
	return _Battle.Contract.GetBattleTimes(&_Battle.CallOpts, _tokenId)
}

// GetBattleTimes is a free data retrieval call binding the contract method 0xf6ac6c71.
//
// Solidity: function getBattleTimes(uint256 _tokenId) view returns(uint256)
func (_Battle *BattleCallerSession) GetBattleTimes(_tokenId *big.Int) (*big.Int, error) {
	return _Battle.Contract.GetBattleTimes(&_Battle.CallOpts, _tokenId)
}

// GetDayBattleTime is a free data retrieval call binding the contract method 0x2efda468.
//
// Solidity: function getDayBattleTime(uint256 _tokenid) view returns(uint256 _times)
func (_Battle *BattleCaller) GetDayBattleTime(opts *bind.CallOpts, _tokenid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getDayBattleTime", _tokenid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDayBattleTime is a free data retrieval call binding the contract method 0x2efda468.
//
// Solidity: function getDayBattleTime(uint256 _tokenid) view returns(uint256 _times)
func (_Battle *BattleSession) GetDayBattleTime(_tokenid *big.Int) (*big.Int, error) {
	return _Battle.Contract.GetDayBattleTime(&_Battle.CallOpts, _tokenid)
}

// GetDayBattleTime is a free data retrieval call binding the contract method 0x2efda468.
//
// Solidity: function getDayBattleTime(uint256 _tokenid) view returns(uint256 _times)
func (_Battle *BattleCallerSession) GetDayBattleTime(_tokenid *big.Int) (*big.Int, error) {
	return _Battle.Contract.GetDayBattleTime(&_Battle.CallOpts, _tokenid)
}

// GetEvolveEggPrice is a free data retrieval call binding the contract method 0x950a8915.
//
// Solidity: function getEvolveEggPrice() view returns(uint256)
func (_Battle *BattleCaller) GetEvolveEggPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getEvolveEggPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEvolveEggPrice is a free data retrieval call binding the contract method 0x950a8915.
//
// Solidity: function getEvolveEggPrice() view returns(uint256)
func (_Battle *BattleSession) GetEvolveEggPrice() (*big.Int, error) {
	return _Battle.Contract.GetEvolveEggPrice(&_Battle.CallOpts)
}

// GetEvolveEggPrice is a free data retrieval call binding the contract method 0x950a8915.
//
// Solidity: function getEvolveEggPrice() view returns(uint256)
func (_Battle *BattleCallerSession) GetEvolveEggPrice() (*big.Int, error) {
	return _Battle.Contract.GetEvolveEggPrice(&_Battle.CallOpts)
}

// GetFightHistorys is a free data retrieval call binding the contract method 0x8fca6a84.
//
// Solidity: function getFightHistorys(address _owner) view returns((uint256,uint8,uint8,uint256,uint256,uint256,uint256)[])
func (_Battle *BattleCaller) GetFightHistorys(opts *bind.CallOpts, _owner common.Address) ([]BattleFightHistory, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getFightHistorys", _owner)

	if err != nil {
		return *new([]BattleFightHistory), err
	}

	out0 := *abi.ConvertType(out[0], new([]BattleFightHistory)).(*[]BattleFightHistory)

	return out0, err

}

// GetFightHistorys is a free data retrieval call binding the contract method 0x8fca6a84.
//
// Solidity: function getFightHistorys(address _owner) view returns((uint256,uint8,uint8,uint256,uint256,uint256,uint256)[])
func (_Battle *BattleSession) GetFightHistorys(_owner common.Address) ([]BattleFightHistory, error) {
	return _Battle.Contract.GetFightHistorys(&_Battle.CallOpts, _owner)
}

// GetFightHistorys is a free data retrieval call binding the contract method 0x8fca6a84.
//
// Solidity: function getFightHistorys(address _owner) view returns((uint256,uint8,uint8,uint256,uint256,uint256,uint256)[])
func (_Battle *BattleCallerSession) GetFightHistorys(_owner common.Address) ([]BattleFightHistory, error) {
	return _Battle.Contract.GetFightHistorys(&_Battle.CallOpts, _owner)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address _sender) view returns(uint256 _reward)
func (_Battle *BattleCaller) GetPendingReward(opts *bind.CallOpts, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getPendingReward", _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address _sender) view returns(uint256 _reward)
func (_Battle *BattleSession) GetPendingReward(_sender common.Address) (*big.Int, error) {
	return _Battle.Contract.GetPendingReward(&_Battle.CallOpts, _sender)
}

// GetPendingReward is a free data retrieval call binding the contract method 0x4df9d6ba.
//
// Solidity: function getPendingReward(address _sender) view returns(uint256 _reward)
func (_Battle *BattleCallerSession) GetPendingReward(_sender common.Address) (*big.Int, error) {
	return _Battle.Contract.GetPendingReward(&_Battle.CallOpts, _sender)
}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _sender) view returns(uint256 _nonce)
func (_Battle *BattleCaller) GetSpendNonce(opts *bind.CallOpts, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getSpendNonce", _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _sender) view returns(uint256 _nonce)
func (_Battle *BattleSession) GetSpendNonce(_sender common.Address) (*big.Int, error) {
	return _Battle.Contract.GetSpendNonce(&_Battle.CallOpts, _sender)
}

// GetSpendNonce is a free data retrieval call binding the contract method 0x241602f9.
//
// Solidity: function getSpendNonce(address _sender) view returns(uint256 _nonce)
func (_Battle *BattleCallerSession) GetSpendNonce(_sender common.Address) (*big.Int, error) {
	return _Battle.Contract.GetSpendNonce(&_Battle.CallOpts, _sender)
}

// GetUserNonce is a free data retrieval call binding the contract method 0x26bb006c.
//
// Solidity: function getUserNonce(address _user, uint256 _tokenid) view returns(uint256 _userNonce)
func (_Battle *BattleCaller) GetUserNonce(opts *bind.CallOpts, _user common.Address, _tokenid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getUserNonce", _user, _tokenid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNonce is a free data retrieval call binding the contract method 0x26bb006c.
//
// Solidity: function getUserNonce(address _user, uint256 _tokenid) view returns(uint256 _userNonce)
func (_Battle *BattleSession) GetUserNonce(_user common.Address, _tokenid *big.Int) (*big.Int, error) {
	return _Battle.Contract.GetUserNonce(&_Battle.CallOpts, _user, _tokenid)
}

// GetUserNonce is a free data retrieval call binding the contract method 0x26bb006c.
//
// Solidity: function getUserNonce(address _user, uint256 _tokenid) view returns(uint256 _userNonce)
func (_Battle *BattleCallerSession) GetUserNonce(_user common.Address, _tokenid *big.Int) (*big.Int, error) {
	return _Battle.Contract.GetUserNonce(&_Battle.CallOpts, _user, _tokenid)
}

// GetZbattle is a free data retrieval call binding the contract method 0x45a23611.
//
// Solidity: function getZbattle() view returns(uint256)
func (_Battle *BattleCaller) GetZbattle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getZbattle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetZbattle is a free data retrieval call binding the contract method 0x45a23611.
//
// Solidity: function getZbattle() view returns(uint256)
func (_Battle *BattleSession) GetZbattle() (*big.Int, error) {
	return _Battle.Contract.GetZbattle(&_Battle.CallOpts)
}

// GetZbattle is a free data retrieval call binding the contract method 0x45a23611.
//
// Solidity: function getZbattle() view returns(uint256)
func (_Battle *BattleCallerSession) GetZbattle() (*big.Int, error) {
	return _Battle.Contract.GetZbattle(&_Battle.CallOpts)
}

// GetxBattle is a free data retrieval call binding the contract method 0xa6d2c617.
//
// Solidity: function getxBattle() view returns(uint256)
func (_Battle *BattleCaller) GetxBattle(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "getxBattle")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetxBattle is a free data retrieval call binding the contract method 0xa6d2c617.
//
// Solidity: function getxBattle() view returns(uint256)
func (_Battle *BattleSession) GetxBattle() (*big.Int, error) {
	return _Battle.Contract.GetxBattle(&_Battle.CallOpts)
}

// GetxBattle is a free data retrieval call binding the contract method 0xa6d2c617.
//
// Solidity: function getxBattle() view returns(uint256)
func (_Battle *BattleCallerSession) GetxBattle() (*big.Int, error) {
	return _Battle.Contract.GetxBattle(&_Battle.CallOpts)
}

// LastTimes is a free data retrieval call binding the contract method 0x4f1e8e52.
//
// Solidity: function lastTimes(address ) view returns(uint256)
func (_Battle *BattleCaller) LastTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "lastTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTimes is a free data retrieval call binding the contract method 0x4f1e8e52.
//
// Solidity: function lastTimes(address ) view returns(uint256)
func (_Battle *BattleSession) LastTimes(arg0 common.Address) (*big.Int, error) {
	return _Battle.Contract.LastTimes(&_Battle.CallOpts, arg0)
}

// LastTimes is a free data retrieval call binding the contract method 0x4f1e8e52.
//
// Solidity: function lastTimes(address ) view returns(uint256)
func (_Battle *BattleCallerSession) LastTimes(arg0 common.Address) (*big.Int, error) {
	return _Battle.Contract.LastTimes(&_Battle.CallOpts, arg0)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Battle *BattleCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Battle *BattleSession) Manager() (common.Address, error) {
	return _Battle.Contract.Manager(&_Battle.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Battle *BattleCallerSession) Manager() (common.Address, error) {
	return _Battle.Contract.Manager(&_Battle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Battle *BattleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Battle *BattleSession) Owner() (common.Address, error) {
	return _Battle.Contract.Owner(&_Battle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Battle *BattleCallerSession) Owner() (common.Address, error) {
	return _Battle.Contract.Owner(&_Battle.CallOpts)
}

// RandomExp is a free data retrieval call binding the contract method 0x44482c11.
//
// Solidity: function randomExp(uint256 _id, uint256 _length, uint256 _value) view returns(uint256)
func (_Battle *BattleCaller) RandomExp(opts *bind.CallOpts, _id *big.Int, _length *big.Int, _value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "randomExp", _id, _length, _value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RandomExp is a free data retrieval call binding the contract method 0x44482c11.
//
// Solidity: function randomExp(uint256 _id, uint256 _length, uint256 _value) view returns(uint256)
func (_Battle *BattleSession) RandomExp(_id *big.Int, _length *big.Int, _value *big.Int) (*big.Int, error) {
	return _Battle.Contract.RandomExp(&_Battle.CallOpts, _id, _length, _value)
}

// RandomExp is a free data retrieval call binding the contract method 0x44482c11.
//
// Solidity: function randomExp(uint256 _id, uint256 _length, uint256 _value) view returns(uint256)
func (_Battle *BattleCallerSession) RandomExp(_id *big.Int, _length *big.Int, _value *big.Int) (*big.Int, error) {
	return _Battle.Contract.RandomExp(&_Battle.CallOpts, _id, _length, _value)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Battle *BattleCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Battle *BattleSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Battle.Contract.Rewards(&_Battle.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Battle *BattleCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Battle.Contract.Rewards(&_Battle.CallOpts, arg0)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Battle *BattleCaller) Signer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "signer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Battle *BattleSession) Signer() (common.Address, error) {
	return _Battle.Contract.Signer(&_Battle.CallOpts)
}

// Signer is a free data retrieval call binding the contract method 0x238ac933.
//
// Solidity: function signer() view returns(address)
func (_Battle *BattleCallerSession) Signer() (common.Address, error) {
	return _Battle.Contract.Signer(&_Battle.CallOpts)
}

// SuplusTimes is a free data retrieval call binding the contract method 0x7796a8e5.
//
// Solidity: function suplusTimes(uint256 _tokenId) view returns(uint256 count, uint256 lasttime)
func (_Battle *BattleCaller) SuplusTimes(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Count    *big.Int
	Lasttime *big.Int
}, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "suplusTimes", _tokenId)

	outstruct := new(struct {
		Count    *big.Int
		Lasttime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Lasttime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SuplusTimes is a free data retrieval call binding the contract method 0x7796a8e5.
//
// Solidity: function suplusTimes(uint256 _tokenId) view returns(uint256 count, uint256 lasttime)
func (_Battle *BattleSession) SuplusTimes(_tokenId *big.Int) (struct {
	Count    *big.Int
	Lasttime *big.Int
}, error) {
	return _Battle.Contract.SuplusTimes(&_Battle.CallOpts, _tokenId)
}

// SuplusTimes is a free data retrieval call binding the contract method 0x7796a8e5.
//
// Solidity: function suplusTimes(uint256 _tokenId) view returns(uint256 count, uint256 lasttime)
func (_Battle *BattleCallerSession) SuplusTimes(_tokenId *big.Int) (struct {
	Count    *big.Int
	Lasttime *big.Int
}, error) {
	return _Battle.Contract.SuplusTimes(&_Battle.CallOpts, _tokenId)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Battle *BattleCaller) Usdt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "usdt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Battle *BattleSession) Usdt() (common.Address, error) {
	return _Battle.Contract.Usdt(&_Battle.CallOpts)
}

// Usdt is a free data retrieval call binding the contract method 0x2f48ab7d.
//
// Solidity: function usdt() view returns(address)
func (_Battle *BattleCallerSession) Usdt() (common.Address, error) {
	return _Battle.Contract.Usdt(&_Battle.CallOpts)
}

// ValidSignature is a free data retrieval call binding the contract method 0xf7a2afc1.
//
// Solidity: function validSignature(uint256 _tokenid, uint256 _exp, uint256 _reward, bytes32 rs, bytes32 ss, uint8 vs) view returns(bool)
func (_Battle *BattleCaller) ValidSignature(opts *bind.CallOpts, _tokenid *big.Int, _exp *big.Int, _reward *big.Int, rs [32]byte, ss [32]byte, vs uint8) (bool, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "validSignature", _tokenid, _exp, _reward, rs, ss, vs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidSignature is a free data retrieval call binding the contract method 0xf7a2afc1.
//
// Solidity: function validSignature(uint256 _tokenid, uint256 _exp, uint256 _reward, bytes32 rs, bytes32 ss, uint8 vs) view returns(bool)
func (_Battle *BattleSession) ValidSignature(_tokenid *big.Int, _exp *big.Int, _reward *big.Int, rs [32]byte, ss [32]byte, vs uint8) (bool, error) {
	return _Battle.Contract.ValidSignature(&_Battle.CallOpts, _tokenid, _exp, _reward, rs, ss, vs)
}

// ValidSignature is a free data retrieval call binding the contract method 0xf7a2afc1.
//
// Solidity: function validSignature(uint256 _tokenid, uint256 _exp, uint256 _reward, bytes32 rs, bytes32 ss, uint8 vs) view returns(bool)
func (_Battle *BattleCallerSession) ValidSignature(_tokenid *big.Int, _exp *big.Int, _reward *big.Int, rs [32]byte, ss [32]byte, vs uint8) (bool, error) {
	return _Battle.Contract.ValidSignature(&_Battle.CallOpts, _tokenid, _exp, _reward, rs, ss, vs)
}

// ZoonerERC20 is a free data retrieval call binding the contract method 0x63a28caa.
//
// Solidity: function zoonerERC20() view returns(address)
func (_Battle *BattleCaller) ZoonerERC20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "zoonerERC20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoonerERC20 is a free data retrieval call binding the contract method 0x63a28caa.
//
// Solidity: function zoonerERC20() view returns(address)
func (_Battle *BattleSession) ZoonerERC20() (common.Address, error) {
	return _Battle.Contract.ZoonerERC20(&_Battle.CallOpts)
}

// ZoonerERC20 is a free data retrieval call binding the contract method 0x63a28caa.
//
// Solidity: function zoonerERC20() view returns(address)
func (_Battle *BattleCallerSession) ZoonerERC20() (common.Address, error) {
	return _Battle.Contract.ZoonerERC20(&_Battle.CallOpts)
}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_Battle *BattleCaller) ZoonerNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Battle.contract.Call(opts, &out, "zoonerNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_Battle *BattleSession) ZoonerNFT() (common.Address, error) {
	return _Battle.Contract.ZoonerNFT(&_Battle.CallOpts)
}

// ZoonerNFT is a free data retrieval call binding the contract method 0xda940945.
//
// Solidity: function zoonerNFT() view returns(address)
func (_Battle *BattleCallerSession) ZoonerNFT() (common.Address, error) {
	return _Battle.Contract.ZoonerNFT(&_Battle.CallOpts)
}

// Userbattle is a paid mutator transaction binding the contract method 0xe20ee7fe.
//
// Solidity: function Userbattle((uint128,uint128,uint256,uint256) _battleinfo, bytes32 r, bytes32 s, uint8 v) returns()
func (_Battle *BattleTransactor) Userbattle(opts *bind.TransactOpts, _battleinfo BattleBattleStruct, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "Userbattle", _battleinfo, r, s, v)
}

// Userbattle is a paid mutator transaction binding the contract method 0xe20ee7fe.
//
// Solidity: function Userbattle((uint128,uint128,uint256,uint256) _battleinfo, bytes32 r, bytes32 s, uint8 v) returns()
func (_Battle *BattleSession) Userbattle(_battleinfo BattleBattleStruct, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Battle.Contract.Userbattle(&_Battle.TransactOpts, _battleinfo, r, s, v)
}

// Userbattle is a paid mutator transaction binding the contract method 0xe20ee7fe.
//
// Solidity: function Userbattle((uint128,uint128,uint256,uint256) _battleinfo, bytes32 r, bytes32 s, uint8 v) returns()
func (_Battle *BattleTransactorSession) Userbattle(_battleinfo BattleBattleStruct, r [32]byte, s [32]byte, v uint8) (*types.Transaction, error) {
	return _Battle.Contract.Userbattle(&_Battle.TransactOpts, _battleinfo, r, s, v)
}

// Battle1 is a paid mutator transaction binding the contract method 0xb7026392.
//
// Solidity: function battle1(uint256 _tokenId, uint8 _monster) returns()
func (_Battle *BattleTransactor) Battle1(opts *bind.TransactOpts, _tokenId *big.Int, _monster uint8) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "battle1", _tokenId, _monster)
}

// Battle1 is a paid mutator transaction binding the contract method 0xb7026392.
//
// Solidity: function battle1(uint256 _tokenId, uint8 _monster) returns()
func (_Battle *BattleSession) Battle1(_tokenId *big.Int, _monster uint8) (*types.Transaction, error) {
	return _Battle.Contract.Battle1(&_Battle.TransactOpts, _tokenId, _monster)
}

// Battle1 is a paid mutator transaction binding the contract method 0xb7026392.
//
// Solidity: function battle1(uint256 _tokenId, uint8 _monster) returns()
func (_Battle *BattleTransactorSession) Battle1(_tokenId *big.Int, _monster uint8) (*types.Transaction, error) {
	return _Battle.Contract.Battle1(&_Battle.TransactOpts, _tokenId, _monster)
}

// EvolveEgg is a paid mutator transaction binding the contract method 0x2851ff81.
//
// Solidity: function evolveEgg(uint256 _tokenId) returns()
func (_Battle *BattleTransactor) EvolveEgg(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "evolveEgg", _tokenId)
}

// EvolveEgg is a paid mutator transaction binding the contract method 0x2851ff81.
//
// Solidity: function evolveEgg(uint256 _tokenId) returns()
func (_Battle *BattleSession) EvolveEgg(_tokenId *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.EvolveEgg(&_Battle.TransactOpts, _tokenId)
}

// EvolveEgg is a paid mutator transaction binding the contract method 0x2851ff81.
//
// Solidity: function evolveEgg(uint256 _tokenId) returns()
func (_Battle *BattleTransactorSession) EvolveEgg(_tokenId *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.EvolveEgg(&_Battle.TransactOpts, _tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Battle *BattleTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Battle *BattleSession) Initialize() (*types.Transaction, error) {
	return _Battle.Contract.Initialize(&_Battle.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Battle *BattleTransactorSession) Initialize() (*types.Transaction, error) {
	return _Battle.Contract.Initialize(&_Battle.TransactOpts)
}

// LayRnd is a paid mutator transaction binding the contract method 0xeb3f2340.
//
// Solidity: function layRnd(uint256 _amount) returns()
func (_Battle *BattleTransactor) LayRnd(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "layRnd", _amount)
}

// LayRnd is a paid mutator transaction binding the contract method 0xeb3f2340.
//
// Solidity: function layRnd(uint256 _amount) returns()
func (_Battle *BattleSession) LayRnd(_amount *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.LayRnd(&_Battle.TransactOpts, _amount)
}

// LayRnd is a paid mutator transaction binding the contract method 0xeb3f2340.
//
// Solidity: function layRnd(uint256 _amount) returns()
func (_Battle *BattleTransactorSession) LayRnd(_amount *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.LayRnd(&_Battle.TransactOpts, _amount)
}

// LayRndDogCoin is a paid mutator transaction binding the contract method 0x03a95d39.
//
// Solidity: function layRndDogCoin(uint256 _amount) returns()
func (_Battle *BattleTransactor) LayRndDogCoin(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "layRndDogCoin", _amount)
}

// LayRndDogCoin is a paid mutator transaction binding the contract method 0x03a95d39.
//
// Solidity: function layRndDogCoin(uint256 _amount) returns()
func (_Battle *BattleSession) LayRndDogCoin(_amount *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.LayRndDogCoin(&_Battle.TransactOpts, _amount)
}

// LayRndDogCoin is a paid mutator transaction binding the contract method 0x03a95d39.
//
// Solidity: function layRndDogCoin(uint256 _amount) returns()
func (_Battle *BattleTransactorSession) LayRndDogCoin(_amount *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.LayRndDogCoin(&_Battle.TransactOpts, _amount)
}

// LayRndUsdt is a paid mutator transaction binding the contract method 0xf97bfc14.
//
// Solidity: function layRndUsdt(uint256 _amount) returns()
func (_Battle *BattleTransactor) LayRndUsdt(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "layRndUsdt", _amount)
}

// LayRndUsdt is a paid mutator transaction binding the contract method 0xf97bfc14.
//
// Solidity: function layRndUsdt(uint256 _amount) returns()
func (_Battle *BattleSession) LayRndUsdt(_amount *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.LayRndUsdt(&_Battle.TransactOpts, _amount)
}

// LayRndUsdt is a paid mutator transaction binding the contract method 0xf97bfc14.
//
// Solidity: function layRndUsdt(uint256 _amount) returns()
func (_Battle *BattleTransactorSession) LayRndUsdt(_amount *big.Int) (*types.Transaction, error) {
	return _Battle.Contract.LayRndUsdt(&_Battle.TransactOpts, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Battle *BattleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Battle *BattleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Battle.Contract.RenounceOwnership(&_Battle.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Battle *BattleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Battle.Contract.RenounceOwnership(&_Battle.TransactOpts)
}

// SetBattleInfo is a paid mutator transaction binding the contract method 0xc32fc5fc.
//
// Solidity: function setBattleInfo(address _manager, address _zoonerERC20, address zoonerNFT) returns()
func (_Battle *BattleTransactor) SetBattleInfo(opts *bind.TransactOpts, _manager common.Address, _zoonerERC20 common.Address, zoonerNFT common.Address) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "setBattleInfo", _manager, _zoonerERC20, zoonerNFT)
}

// SetBattleInfo is a paid mutator transaction binding the contract method 0xc32fc5fc.
//
// Solidity: function setBattleInfo(address _manager, address _zoonerERC20, address zoonerNFT) returns()
func (_Battle *BattleSession) SetBattleInfo(_manager common.Address, _zoonerERC20 common.Address, zoonerNFT common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SetBattleInfo(&_Battle.TransactOpts, _manager, _zoonerERC20, zoonerNFT)
}

// SetBattleInfo is a paid mutator transaction binding the contract method 0xc32fc5fc.
//
// Solidity: function setBattleInfo(address _manager, address _zoonerERC20, address zoonerNFT) returns()
func (_Battle *BattleTransactorSession) SetBattleInfo(_manager common.Address, _zoonerERC20 common.Address, zoonerNFT common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SetBattleInfo(&_Battle.TransactOpts, _manager, _zoonerERC20, zoonerNFT)
}

// SetBlack is a paid mutator transaction binding the contract method 0x98612dde.
//
// Solidity: function setBlack(address[] _user) returns()
func (_Battle *BattleTransactor) SetBlack(opts *bind.TransactOpts, _user []common.Address) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "setBlack", _user)
}

// SetBlack is a paid mutator transaction binding the contract method 0x98612dde.
//
// Solidity: function setBlack(address[] _user) returns()
func (_Battle *BattleSession) SetBlack(_user []common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SetBlack(&_Battle.TransactOpts, _user)
}

// SetBlack is a paid mutator transaction binding the contract method 0x98612dde.
//
// Solidity: function setBlack(address[] _user) returns()
func (_Battle *BattleTransactorSession) SetBlack(_user []common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SetBlack(&_Battle.TransactOpts, _user)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x7f1d259b.
//
// Solidity: function setTokenAddr(address _usdt, address _dogCoin) returns()
func (_Battle *BattleTransactor) SetTokenAddr(opts *bind.TransactOpts, _usdt common.Address, _dogCoin common.Address) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "setTokenAddr", _usdt, _dogCoin)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x7f1d259b.
//
// Solidity: function setTokenAddr(address _usdt, address _dogCoin) returns()
func (_Battle *BattleSession) SetTokenAddr(_usdt common.Address, _dogCoin common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SetTokenAddr(&_Battle.TransactOpts, _usdt, _dogCoin)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x7f1d259b.
//
// Solidity: function setTokenAddr(address _usdt, address _dogCoin) returns()
func (_Battle *BattleTransactorSession) SetTokenAddr(_usdt common.Address, _dogCoin common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SetTokenAddr(&_Battle.TransactOpts, _usdt, _dogCoin)
}

// SetTokenids is a paid mutator transaction binding the contract method 0x76693449.
//
// Solidity: function setTokenids(uint256[] _tokenids) returns()
func (_Battle *BattleTransactor) SetTokenids(opts *bind.TransactOpts, _tokenids []*big.Int) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "setTokenids", _tokenids)
}

// SetTokenids is a paid mutator transaction binding the contract method 0x76693449.
//
// Solidity: function setTokenids(uint256[] _tokenids) returns()
func (_Battle *BattleSession) SetTokenids(_tokenids []*big.Int) (*types.Transaction, error) {
	return _Battle.Contract.SetTokenids(&_Battle.TransactOpts, _tokenids)
}

// SetTokenids is a paid mutator transaction binding the contract method 0x76693449.
//
// Solidity: function setTokenids(uint256[] _tokenids) returns()
func (_Battle *BattleTransactorSession) SetTokenids(_tokenids []*big.Int) (*types.Transaction, error) {
	return _Battle.Contract.SetTokenids(&_Battle.TransactOpts, _tokenids)
}

// SignAddr is a paid mutator transaction binding the contract method 0x5e5ecc63.
//
// Solidity: function signAddr(address _sign) returns()
func (_Battle *BattleTransactor) SignAddr(opts *bind.TransactOpts, _sign common.Address) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "signAddr", _sign)
}

// SignAddr is a paid mutator transaction binding the contract method 0x5e5ecc63.
//
// Solidity: function signAddr(address _sign) returns()
func (_Battle *BattleSession) SignAddr(_sign common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SignAddr(&_Battle.TransactOpts, _sign)
}

// SignAddr is a paid mutator transaction binding the contract method 0x5e5ecc63.
//
// Solidity: function signAddr(address _sign) returns()
func (_Battle *BattleTransactorSession) SignAddr(_sign common.Address) (*types.Transaction, error) {
	return _Battle.Contract.SignAddr(&_Battle.TransactOpts, _sign)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Battle *BattleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Battle *BattleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Battle.Contract.TransferOwnership(&_Battle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Battle *BattleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Battle.Contract.TransferOwnership(&_Battle.TransactOpts, newOwner)
}

// WithDrawRewards is a paid mutator transaction binding the contract method 0x696c81b1.
//
// Solidity: function withDrawRewards() returns()
func (_Battle *BattleTransactor) WithDrawRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Battle.contract.Transact(opts, "withDrawRewards")
}

// WithDrawRewards is a paid mutator transaction binding the contract method 0x696c81b1.
//
// Solidity: function withDrawRewards() returns()
func (_Battle *BattleSession) WithDrawRewards() (*types.Transaction, error) {
	return _Battle.Contract.WithDrawRewards(&_Battle.TransactOpts)
}

// WithDrawRewards is a paid mutator transaction binding the contract method 0x696c81b1.
//
// Solidity: function withDrawRewards() returns()
func (_Battle *BattleTransactorSession) WithDrawRewards() (*types.Transaction, error) {
	return _Battle.Contract.WithDrawRewards(&_Battle.TransactOpts)
}

// BattleBattleEventIterator is returned from FilterBattleEvent and is used to iterate over the raw logs and unpacked data for BattleEvent events raised by the Battle contract.
type BattleBattleEventIterator struct {
	Event *BattleBattleEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleBattleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleBattleEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleBattleEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleBattleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleBattleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleBattleEvent represents a BattleEvent event raised by the Battle contract.
type BattleBattleEvent struct {
	TokenId *big.Int
	Monster uint8
	User    common.Address
	Result  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBattleEvent is a free log retrieval operation binding the contract event 0xa10960c01512bbda4e4d2dfc1b3452000df207c8eccb9097485cbcdd078c9af1.
//
// Solidity: event BattleEvent(uint256 indexed _tokenId, uint8 monster, address user, uint8 result)
func (_Battle *BattleFilterer) FilterBattleEvent(opts *bind.FilterOpts, _tokenId []*big.Int) (*BattleBattleEventIterator, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "BattleEvent", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BattleBattleEventIterator{contract: _Battle.contract, event: "BattleEvent", logs: logs, sub: sub}, nil
}

// WatchBattleEvent is a free log subscription operation binding the contract event 0xa10960c01512bbda4e4d2dfc1b3452000df207c8eccb9097485cbcdd078c9af1.
//
// Solidity: event BattleEvent(uint256 indexed _tokenId, uint8 monster, address user, uint8 result)
func (_Battle *BattleFilterer) WatchBattleEvent(opts *bind.WatchOpts, sink chan<- *BattleBattleEvent, _tokenId []*big.Int) (event.Subscription, error) {

	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "BattleEvent", _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleBattleEvent)
				if err := _Battle.contract.UnpackLog(event, "BattleEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBattleEvent is a log parse operation binding the contract event 0xa10960c01512bbda4e4d2dfc1b3452000df207c8eccb9097485cbcdd078c9af1.
//
// Solidity: event BattleEvent(uint256 indexed _tokenId, uint8 monster, address user, uint8 result)
func (_Battle *BattleFilterer) ParseBattleEvent(log types.Log) (*BattleBattleEvent, error) {
	event := new(BattleBattleEvent)
	if err := _Battle.contract.UnpackLog(event, "BattleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Battle contract.
type BattleEmergencyWithdrawIterator struct {
	Event *BattleEmergencyWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleEmergencyWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleEmergencyWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleEmergencyWithdraw represents a EmergencyWithdraw event raised by the Battle contract.
type BattleEmergencyWithdraw struct {
	User    common.Address
	Pid     *big.Int
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 _tokenId)
func (_Battle *BattleFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, user []common.Address, pid []*big.Int) (*BattleEmergencyWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return &BattleEmergencyWithdrawIterator{contract: _Battle.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 _tokenId)
func (_Battle *BattleFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *BattleEmergencyWithdraw, user []common.Address, pid []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var pidRule []interface{}
	for _, pidItem := range pid {
		pidRule = append(pidRule, pidItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "EmergencyWithdraw", userRule, pidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleEmergencyWithdraw)
				if err := _Battle.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0xbb757047c2b5f3974fe26b7c10f732e7bce710b0952a71082702781e62ae0595.
//
// Solidity: event EmergencyWithdraw(address indexed user, uint256 indexed pid, uint256 _tokenId)
func (_Battle *BattleFilterer) ParseEmergencyWithdraw(log types.Log) (*BattleEmergencyWithdraw, error) {
	event := new(BattleEmergencyWithdraw)
	if err := _Battle.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleEvolveIterator is returned from FilterEvolve and is used to iterate over the raw logs and unpacked data for Evolve events raised by the Battle contract.
type BattleEvolveIterator struct {
	Event *BattleEvolve // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleEvolveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleEvolve)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleEvolve)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleEvolveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleEvolveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleEvolve represents a Evolve event raised by the Battle contract.
type BattleEvolve struct {
	TokenId *big.Int
	Owner   common.Address
	Dna     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEvolve is a free log retrieval operation binding the contract event 0x0bb0e4d988fc9cd4f6e98f64ff90bce64b8f322fc1ffa6e7ac88691712b6c379.
//
// Solidity: event Evolve(uint256 indexed tokenId, address owner, uint256 dna)
func (_Battle *BattleFilterer) FilterEvolve(opts *bind.FilterOpts, tokenId []*big.Int) (*BattleEvolveIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "Evolve", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BattleEvolveIterator{contract: _Battle.contract, event: "Evolve", logs: logs, sub: sub}, nil
}

// WatchEvolve is a free log subscription operation binding the contract event 0x0bb0e4d988fc9cd4f6e98f64ff90bce64b8f322fc1ffa6e7ac88691712b6c379.
//
// Solidity: event Evolve(uint256 indexed tokenId, address owner, uint256 dna)
func (_Battle *BattleFilterer) WatchEvolve(opts *bind.WatchOpts, sink chan<- *BattleEvolve, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "Evolve", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleEvolve)
				if err := _Battle.contract.UnpackLog(event, "Evolve", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEvolve is a log parse operation binding the contract event 0x0bb0e4d988fc9cd4f6e98f64ff90bce64b8f322fc1ffa6e7ac88691712b6c379.
//
// Solidity: event Evolve(uint256 indexed tokenId, address owner, uint256 dna)
func (_Battle *BattleFilterer) ParseEvolve(log types.Log) (*BattleEvolve, error) {
	event := new(BattleEvolve)
	if err := _Battle.contract.UnpackLog(event, "Evolve", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Battle contract.
type BattleInitializedIterator struct {
	Event *BattleInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleInitialized represents a Initialized event raised by the Battle contract.
type BattleInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Battle *BattleFilterer) FilterInitialized(opts *bind.FilterOpts) (*BattleInitializedIterator, error) {

	logs, sub, err := _Battle.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BattleInitializedIterator{contract: _Battle.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Battle *BattleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BattleInitialized) (event.Subscription, error) {

	logs, sub, err := _Battle.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleInitialized)
				if err := _Battle.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Battle *BattleFilterer) ParseInitialized(log types.Log) (*BattleInitialized, error) {
	event := new(BattleInitialized)
	if err := _Battle.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Battle contract.
type BattleOwnershipTransferredIterator struct {
	Event *BattleOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleOwnershipTransferred represents a OwnershipTransferred event raised by the Battle contract.
type BattleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Battle *BattleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BattleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BattleOwnershipTransferredIterator{contract: _Battle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Battle *BattleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BattleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleOwnershipTransferred)
				if err := _Battle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Battle *BattleFilterer) ParseOwnershipTransferred(log types.Log) (*BattleOwnershipTransferred, error) {
	event := new(BattleOwnershipTransferred)
	if err := _Battle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleSpawnIterator is returned from FilterSpawn and is used to iterate over the raw logs and unpacked data for Spawn events raised by the Battle contract.
type BattleSpawnIterator struct {
	Event *BattleSpawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleSpawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleSpawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleSpawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleSpawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleSpawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleSpawn represents a Spawn event raised by the Battle contract.
type BattleSpawn struct {
	Buyer  common.Address
	Amount *big.Int
	Tribe  uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSpawn is a free log retrieval operation binding the contract event 0x2c8b4980081f56afaef1c27ceea6d0897119cbdb0690cdc56a1967a68ec3afcb.
//
// Solidity: event Spawn(address indexed buyer, uint256 amount, uint8 _tribe)
func (_Battle *BattleFilterer) FilterSpawn(opts *bind.FilterOpts, buyer []common.Address) (*BattleSpawnIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "Spawn", buyerRule)
	if err != nil {
		return nil, err
	}
	return &BattleSpawnIterator{contract: _Battle.contract, event: "Spawn", logs: logs, sub: sub}, nil
}

// WatchSpawn is a free log subscription operation binding the contract event 0x2c8b4980081f56afaef1c27ceea6d0897119cbdb0690cdc56a1967a68ec3afcb.
//
// Solidity: event Spawn(address indexed buyer, uint256 amount, uint8 _tribe)
func (_Battle *BattleFilterer) WatchSpawn(opts *bind.WatchOpts, sink chan<- *BattleSpawn, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "Spawn", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleSpawn)
				if err := _Battle.contract.UnpackLog(event, "Spawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSpawn is a log parse operation binding the contract event 0x2c8b4980081f56afaef1c27ceea6d0897119cbdb0690cdc56a1967a68ec3afcb.
//
// Solidity: event Spawn(address indexed buyer, uint256 amount, uint8 _tribe)
func (_Battle *BattleFilterer) ParseSpawn(log types.Log) (*BattleSpawn, error) {
	event := new(BattleSpawn)
	if err := _Battle.contract.UnpackLog(event, "Spawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleUserBattleEventIterator is returned from FilterUserBattleEvent and is used to iterate over the raw logs and unpacked data for UserBattleEvent events raised by the Battle contract.
type BattleUserBattleEventIterator struct {
	Event *BattleUserBattleEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleUserBattleEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleUserBattleEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleUserBattleEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleUserBattleEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleUserBattleEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleUserBattleEvent represents a UserBattleEvent event raised by the Battle contract.
type BattleUserBattleEvent struct {
	Id          *big.Int
	BlockNumber *big.Int
	BlockTime   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserBattleEvent is a free log retrieval operation binding the contract event 0xa5a69883071d79e0c4978c7b9b1b19512385f4a15826cf8463893c39734dca37.
//
// Solidity: event UserBattleEvent(uint256 indexed Id, uint256 blockNumber, uint256 blockTime)
func (_Battle *BattleFilterer) FilterUserBattleEvent(opts *bind.FilterOpts, Id []*big.Int) (*BattleUserBattleEventIterator, error) {

	var IdRule []interface{}
	for _, IdItem := range Id {
		IdRule = append(IdRule, IdItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "UserBattleEvent", IdRule)
	if err != nil {
		return nil, err
	}
	return &BattleUserBattleEventIterator{contract: _Battle.contract, event: "UserBattleEvent", logs: logs, sub: sub}, nil
}

// WatchUserBattleEvent is a free log subscription operation binding the contract event 0xa5a69883071d79e0c4978c7b9b1b19512385f4a15826cf8463893c39734dca37.
//
// Solidity: event UserBattleEvent(uint256 indexed Id, uint256 blockNumber, uint256 blockTime)
func (_Battle *BattleFilterer) WatchUserBattleEvent(opts *bind.WatchOpts, sink chan<- *BattleUserBattleEvent, Id []*big.Int) (event.Subscription, error) {

	var IdRule []interface{}
	for _, IdItem := range Id {
		IdRule = append(IdRule, IdItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "UserBattleEvent", IdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleUserBattleEvent)
				if err := _Battle.contract.UnpackLog(event, "UserBattleEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserBattleEvent is a log parse operation binding the contract event 0xa5a69883071d79e0c4978c7b9b1b19512385f4a15826cf8463893c39734dca37.
//
// Solidity: event UserBattleEvent(uint256 indexed Id, uint256 blockNumber, uint256 blockTime)
func (_Battle *BattleFilterer) ParseUserBattleEvent(log types.Log) (*BattleUserBattleEvent, error) {
	event := new(BattleUserBattleEvent)
	if err := _Battle.contract.UnpackLog(event, "UserBattleEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BattleWithdrawEventIterator is returned from FilterWithdrawEvent and is used to iterate over the raw logs and unpacked data for WithdrawEvent events raised by the Battle contract.
type BattleWithdrawEventIterator struct {
	Event *BattleWithdrawEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BattleWithdrawEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BattleWithdrawEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BattleWithdrawEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BattleWithdrawEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BattleWithdrawEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BattleWithdrawEvent represents a WithdrawEvent event raised by the Battle contract.
type BattleWithdrawEvent struct {
	User   common.Address
	Round  *big.Int
	Amound *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEvent is a free log retrieval operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed user, uint256 indexed _round, uint256 _amound)
func (_Battle *BattleFilterer) FilterWithdrawEvent(opts *bind.FilterOpts, user []common.Address, _round []*big.Int) (*BattleWithdrawEventIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}

	logs, sub, err := _Battle.contract.FilterLogs(opts, "WithdrawEvent", userRule, _roundRule)
	if err != nil {
		return nil, err
	}
	return &BattleWithdrawEventIterator{contract: _Battle.contract, event: "WithdrawEvent", logs: logs, sub: sub}, nil
}

// WatchWithdrawEvent is a free log subscription operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed user, uint256 indexed _round, uint256 _amound)
func (_Battle *BattleFilterer) WatchWithdrawEvent(opts *bind.WatchOpts, sink chan<- *BattleWithdrawEvent, user []common.Address, _round []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var _roundRule []interface{}
	for _, _roundItem := range _round {
		_roundRule = append(_roundRule, _roundItem)
	}

	logs, sub, err := _Battle.contract.WatchLogs(opts, "WithdrawEvent", userRule, _roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BattleWithdrawEvent)
				if err := _Battle.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawEvent is a log parse operation binding the contract event 0x5bb95829671915ece371da722f91d5371159095dcabf2f75cd6c53facb7e1bab.
//
// Solidity: event WithdrawEvent(address indexed user, uint256 indexed _round, uint256 _amound)
func (_Battle *BattleFilterer) ParseWithdrawEvent(log types.Log) (*BattleWithdrawEvent, error) {
	event := new(BattleWithdrawEvent)
	if err := _Battle.contract.UnpackLog(event, "WithdrawEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
