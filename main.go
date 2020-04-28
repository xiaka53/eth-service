package main

import (
	"github.com/xiaka53/DeployAndLog/lib"
	"github.com/xiaka53/eth-service/controller/task"
	"github.com/xiaka53/eth-service/public"
	"github.com/xiaka53/eth-service/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title ETH_SERVER
// @version 1.0
// @description 以太坊功能服务
// @contact.name 陶然
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	if err := lib.InitModule("./conf/dev/", []string{"base", "mysql"}); err != nil {
		log.Fatal(err)
	}
	if err := public.InitMysql(); err != nil {
		log.Fatal(err)
	}
	{
		if err := task.BlogRefreshInit(); err != nil {
			log.Fatal(err)
		}
	}
	public.InitValidate()
	defer lib.Destroy()
	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	router.HttpServerStop()
}
