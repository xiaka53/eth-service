package dto

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/xiaka53/eth-service/public"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type GetBalanceValidateInput struct {
	Address string `form:"address" json:"address" validate:"eth_addr"`
}

func (o *GetBalanceValidateInput) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("en")
	}
	if err = public.Validate.Struct(o); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return
}

type SendValidateInput struct {
	From  string  `form:"from" json:"from" validate:"eth_addr"`
	To    string  `form:"to" json:"to" validate:"eth_addr"`
	Value float64 `form:"value" json:"value" validate:"min=0"`
	Gas   float64 `form:"gas" json:"gas" validate:"min=0"`
}

func (o *SendValidateInput) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("en")
	}
	if err = public.Validate.Struct(o); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return
}

type TransferValidateInput struct {
	Address string `form:"address" json:"from" validate:"eth_addr"`
	From    int    `form:"from" json:"from" validate:"gte=1"`
	Amount  int    `form:"amount" json:"amount" validate:"gte=1"`
}

func (o *TransferValidateInput) BindingValidParams(c *gin.Context) (err error) {
	if err = c.ShouldBind(o); err != nil {
		return
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = public.Uni.GetTranslator("en")
	}
	if err = public.Validate.Struct(o); err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return
}
