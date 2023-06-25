package models

import "time"

type EthDomain struct {
	Id         int64     `orm:"pk;auto;description(主键id)"  form:"id" json:"id" `
	DecodeData string    `orm:"size(2048);description(数据)" form:"decode_data" json:"decode_data"`
	Creator    string    `orm:"size(42);description(发送者地址)" form:"creator" json:"creator"`
	Owner      string    `orm:"size(42);description(接受者地址)" form:"owner" json:"owner"`
	Hash       string    `orm:"description(hash)" form:"hash" json:"hash"`
	Content    string    `orm:"size(255);description(内容)" form:"content" json:"content"`
	Type       string    `orm:"size(10);description(类型)" form:"type" json:"type"`
	BlockTime  uint64    `orm:"description(时间)" form:"block_time" json:"block_time"`
	Ctime      time.Time `orm:"column(ctime);type(datetime)"`
	Classify   int64     `orm:"description(类型)" form:"classify" json:"classify"`
}

func (a *EthDomain) TableName() string {
	return EthDomainTBName()
}

// 多字段索引
func (u *EthDomain) TableIndex() [][]string {
	return [][]string{
		[]string{"id"},
	}
}
