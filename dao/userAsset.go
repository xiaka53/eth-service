package dao

import (
	"fmt"
	"github.com/xiaka53/eth-service/public"
)

type UserAsset struct {
	UserID          int     `json:"user_id" gorm:"cloumn(user_id)" description:"用户ID"`
	Coin            string  `json:"coin" gorm:"cloumn(coin)" description:"币种名称"`
	Asset           float64 `json:"asset" gorm:"cloumn(asset)" description:"可用资产"`
	ForzenAsset     float64 `json:"forzen_asset" gorm:"cloumn(forzen_asset)" description:"冻结资产"`
	MinerAsset      float64 `json:"miner_asset" gorm:"cloumn(miner_asset)" description:"矿池收益"`
	DynamicAsset    string  `json:"dynamic_asset" gorm:"cloumn(dynamic_asset)" description:"动态收益"`
	Address         string  `json:"address" gorm:"cloumn(address)" description:"地址"`
	TotalRecharge   string  `json:"total_recharge" gorm:"cloumn(total_recharge)" description:"总充值金额(USD)"`
	WithdrawAddress string  `json:"withdraw_address" gorm:"cloumn(withdraw_address)" description:"提币地址"`
	WaitClear       float64 `json:"wait_clear" gorm:"cloumn(wait_clear)" description:"等待清除数据"`
	CreateTime      string  `json:"createtime" gorm:"cloumn(createtime)" description:"创建时间"`
}

func (u *UserAsset) TableName() string {
	return "user_asset"
}

func (u *UserAsset) Updates() {
	coin := (&Coin{Token: u.Coin}).GetName()
	public.HzcPool.Exec(fmt.Sprintf("update %s set coin=%v-forzen_asset where address=%s and coin=%s", u.TableName(), u.Asset, u.Address, coin))
}
