package dao

import (
	"github.com/xiaka53/eth-service/public"
	"time"
)

type Hash struct {
	Id         int     `json:"id" gorm:"cloumn(id);primary_key" description:"ID"`
	BlokNum    int     `json:"blok_num" gorm:"cloumn(blok_num)" description:"块搞"`
	Hash       string  `json:"hash" gorm:"cloumn(hash)" description:"交易hash（不是块hash）"`
	Send       string  `json:"send" gorm:"cloumn(send)" description:"转出地址"`
	To         string  `json:"to" gorm:"cloumn(to)" description:"收入地址"`
	Token      string  `json:"token" gorm:"cloumn(token)" description:"代币token"`
	TokenName  string  `json:"token_name" gorm:"cloumn(token_name)" description:"代币名称"`
	Num        float64 `json:"num" gorm:"cloumn(num)" description:"交易数量"`
	Gas        string  `json:"gas" gorm:"cloumn(gas)" description:"手续费"`
	GasPrice   string  `json:"gas_price" gorm:"cloumn(gas_price)" description:"手续费单价"`
	Createtime int     `json:"createtime" gorm:"cloumn(createtime)" description:"创建时间"`
}

func (h *Hash) TableName() string {
	return "hash"
}

func (h *Hash) Create() error {
	h.Createtime = int(time.Now().Unix())
	return public.SqlPool.Create(h).Error
}

func (h *Hash) First(param string) error {
	return public.SqlPool.Where(h).Order("blok_num desc").Select(param).First(h).Error
}

func (h *Hash) Transfer(from, amort int) []public.Transfer {
	var data []public.Transfer
	public.SqlPool.Table(h.TableName()).Where("send=?", h.Send).Or("to=?", h.Send).Order("createtime desc").Limit(amort).Offset(from).Find(&data)
	return data
}
