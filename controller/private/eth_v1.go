package private

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaka53/eth-service/controller/task"
	"github.com/xiaka53/eth-service/dao"
	"github.com/xiaka53/eth-service/dto"
	"github.com/xiaka53/eth-service/middleware"
	"github.com/xiaka53/eth-service/public"
)

type v1 struct {
	public.Eth
}

func EthV1(router *gin.RouterGroup) {
	v := &v1{}
	v.Eth = dao.NewEth()
	router.GET("new_address", v.getAddress)
	router.GET("get_balance", v.getBalance)
	router.POST("send", v.send)
	router.GET("transfer", v.transfer)
	router.GET("estimate_gas", v.estimate_gas)
}

// @Summary 获取新地址
// @Tags 地址类
// @Id 001
// @Produce  json
// @Success 200 {string} string
// @Router /v1/new_address [get]
func (v *v1) getAddress(c *gin.Context) {
	address := task.GetNewAddress()
	middleware.ResponseSuccess(c, address)
}

// @Summary 获取余额
// @Tags 地址类
// @Id 002
// @Produce  json
// @Param address query string true "地址"
// @Success 200 {string} string
// @Router /v1/get_balance [get]
func (v *v1) getBalance(c *gin.Context) {
	var (
		param dto.GetBalanceValidateInput
		err   error
		value float64
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	if value, err = v.Eth.GetBalance(param.Address); err != nil {
		middleware.ResponseError(c, middleware.ErrorCode, middleware.PromptArr[middleware.ErrorCode])
		return
	}
	middleware.ResponseSuccess(c, value)
	return
}

// @Summary 交易
// @Tags 地址类
// @Id 003
// @Produce  json
// @Param from query string true "转账地址"
// @Param to query string true "收款地址"
// @Param value query string true "转出数量"
// @Param gas query string true "手续费"
// @Success 200 {string} string
// @Router /v1/send [post]
func (v *v1) send(c *gin.Context) {
	var (
		param dto.SendValidateInput
		hash  string
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	if hash, err = v.Eth.Send(param.From, param.To, param.Value, param.Gas); err != nil {
		middleware.ResponseError(c, middleware.ErrorCode, middleware.PromptArr[middleware.ErrorCode])
		return
	}
	middleware.ResponseSuccess(c, hash)
}

// @Summary 交易列表
// @Tags 地址类
// @Id 004
// @Produce  json
// @Param address query string true "地址"
// @Param amount query string true "每页数量"
// @Param from query string true "第几页"
// @Success 200 {string} string
// @Router /v1/transfer [get]
func (v *v1) transfer(c *gin.Context) {
	var (
		param dto.TransferValidateInput
		data  []public.Transfer
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	data = v.Eth.Transfer(param.Address, param.Amount, param.From)
	middleware.ResponseSuccess(c, data)
}

// @Summary 获取手续费
// @Tags 地址类
// @Id 005
// @Produce  json
// @Param from query string true "转账地址"
// @Param to query string true "收款地址"
// @Param value query string true "转出数量"
// @Success 200 {string} string
// @Router /v1/estimate_gas [get]
func (v *v1) estimate_gas(c *gin.Context) {
	var (
		param dto.SendValidateInput
		gas   float64
		err   error
	)
	if err = (&param).BindingValidParams(c); err != nil {
		middleware.ResponseError(c, middleware.ParameterError, err)
		return
	}
	gas = v.Eth.EstimateGas(param.From, param.To, param.Value)
	middleware.ResponseSuccess(c, gas)
}
