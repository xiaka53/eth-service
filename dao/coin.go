package dao

import "github.com/xiaka53/eth-service/public"

type Coin struct {
	Id      int     `json:"id" gorm:"cloumn(id);primary_key" description:"用户ID"`
	Name    string  `json:"name" gorm:"cloumn(name)" description:"币种名称"`
	MinHold float64 `json:"min_hold" gorm:"cloumn(min_hold)" description:"最少持币"`
	Token   string  `json:"token" gorm:"cloumn(token)" description:"币种token"`
	IsPool  string  `json:"is_pool" gorm:"cloumn(is_pool)" description:"是否开启矿池Y->开启｜N->关闭"`
	Status  string  `json:"status" gorm:"cloumn(status)" description:"是否使用Y->使用｜N->关闭"`
}

func (c *Coin) TableName() string {
	return "coin"
}

func (c *Coin) GetName() string {
	c.Status = "Y"
	public.HzcPool.Where(c).First(c)
	if len(c.Name) < 1 {
		return "HZC"
	}
	return c.Name
}
