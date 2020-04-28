package dao

import (
	"github.com/xiaka53/eth-service/public"
	"time"
)

type Address struct {
	Address    string `json:"address"  gorm:"cloumn(address);primary_key" description:"地址"`
	Coin       string `json:"coin"  gorm:"cloumn(coin)" description:"使用用途：eth｜usdt"`
	Status     string `json:"status"  gorm:"cloumn(status)" description:"使用状态Y->已使用｜N->未使用"`
	Createtime int    `json:"createtime"  gorm:"cloumn(createtime)" description:"创建时间"`
	Updatetime int    `json:"updatetime"  gorm:"cloumn(updatetime)" description:"使用时间"`
}

func (a *Address) TableName() string {
	return "address"
}

func (a *Address) Find() map[string]bool {
	var data []Address
	list := make(map[string]bool)
	public.SqlPool.Where(a).Find(&data)
	for _, v := range data {
		list[v.Address] = true
	}
	return list
}

func (a *Address) AddressCount() int {
	var count int
	public.SqlPool.Where("status='Y").Count(&count)
	return count
}

func (a *Address) Create() {
	a.Createtime = int(time.Now().Unix())
	public.SqlPool.Create(a)
}

func (a *Address) Update() error {
	newA := *a
	newA.Updatetime = int(time.Now().Unix())
	newA.Status = "Y"
	return public.SqlPool.Table(a.TableName()).Where(a).Updates(newA).Error
}
