package public

import (
	"github.com/e421083458/gorm"
	"github.com/xiaka53/DeployAndLog/lib"
)

var (
	SqlPool *gorm.DB
	HzcPool *gorm.DB
)

//数据库初始化
func InitMysql() (err error) {
	if SqlPool, err = lib.GetGormPool("default"); err != nil {
		return
	}
	if HzcPool, err = lib.GetGormPool("hzc"); err != nil {
		return
	}
	return
}
