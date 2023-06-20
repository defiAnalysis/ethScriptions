package models

type EthScription struct {
	Id         int64  `orm:"pk;auto;description(主键id)"  form:"id" json:"id" `
	DecodeData string `orm:"size(2048);description(数据)" form:"decode_data" json:"decode_data"`
	Creator    string `orm:"size(42);description(发送者地址)" form:"creator" json:"creator"`
	Owner      string `orm:"size(42);description(接受者地址)" form:"owner" json:"owner"`
	Hash       string `orm:"description(hash)" form:"hash" json:"hash"`
	BlockTime  uint64 `orm:"description(时间)" form:"block_time" json:"block_time"`
}

func (a *EthScription) TableName() string {
	return EthScriptionTBName()
}

// 多字段索引
func (u *EthScription) TableIndex() [][]string {
	return [][]string{
		[]string{"id"},
	}
}
