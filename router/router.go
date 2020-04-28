package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/xiaka53/eth-service/controller/private"
	_ "github.com/xiaka53/eth-service/docs"
	"github.com/xiaka53/eth-service/middleware"
)

//路由初始化
func InitRouter(middlewares ...gin.HandlerFunc) (router *gin.Engine) {
	router = gin.Default()
	router.Use(middlewares...)
	router.Use(middleware.IPAuthMiddleware(), middleware.RecoverMiddleware(), middleware.RequestLog())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//版本：v1
	v1 := router.Group("v1")
	private.EthV1(v1)
	return
}
