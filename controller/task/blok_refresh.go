package task

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/xiaka53/eth-service/dao"
	"github.com/xiaka53/eth-service/public"
	"math"
	"math/big"
	"time"
)

type blogRefresh struct {
	Chan   chan int
	Number int
}

type result struct {
	Transactions []transactions `json:"transactions"`
}

type transactions struct {
	From     string `json:"from"`
	Hash     string `json:"hash"`
	To       string `json:"to"`
	Value    string `json:"value"`
	Gas      string `json:"gas"`
	GasPrice string `json:"gasPrice"`
	Input    string `json:"input"`
}

func (b *blogRefresh) setBlock(block int) {
	go func() {
		b.Chan <- block
	}()
}

func (b *blogRefresh) getBlock() int {
	return <-b.Chan
}

func (b *blogRefresh) blockServer() {
	t := time.NewTicker(1 * time.Second / 10)
	for {
		block := b.getBlockNum()
		if block > b.Number {
			b.Number++
			b.setBlock(b.Number)
		}
		<-t.C
	}
}

func (b *blogRefresh) getBlockNum() int {
	var result string
	client := public.GetClient()
	defer client.Close()
	if err := client.Call(&result, "eth_blockNumber"); err != nil {
		fmt.Println(err)
	}
	num, _ := new(big.Int).SetString(result[2:], 16)
	nums := num.Int64()
	return int(nums)
}

func (b *blogRefresh) hashServer() {
	for {
		block := b.getBlock()
		client := public.GetClient()
		defer client.Close()
		var result result
		blockHash := hexutil.EncodeUint64(uint64(block))
		_ = client.Call(&result, "eth_getBlockByNumber", blockHash, true)
		for _, v := range result.Transactions {
			if _, ok := EthAddress.Address[v.From]; ok {
				goto Send
			}
			if _, ok := EthAddress.Address[v.To]; ok {
				goto Send
			}
			continue
		Send:
			num := new(big.Float)
			num, _ = num.SetString(v.Value)
			nums, _ := num.Float64()
			data := &dao.Hash{
				BlokNum:  block,
				Hash:     v.Hash,
				Send:     v.From,
				To:       v.To,
				Num:      nums / math.Pow(10, 18),
				Gas:      v.Gas,
				GasPrice: v.GasPrice,
			}
			if data.Num == 0 && v.Input[:10] == "0xa9059cbb" {
				data.Token = data.To
				data.To, data.Num = b.inputData(v.Input)
			}
			if err := data.Create(); err != nil {
				fmt.Println(err)
			}
			continue
		}
		if err := (&dao.Blog{}).Update(block); err != nil {
			fmt.Println(err)
		}
	}
}

func (b *blogRefresh) inputData(to string) (string, float64) {
	toAddress := to[10:64]
	if toAddress[:2] != "0x" && toAddress[2:] != "0X" {
		toAddress = "0x" + toAddress
	}
	toAddress = toAddress[24:]
	valus := to[74:]
	if len(valus) > 2 {
		if valus[:2] == "0x" || valus[2:] == "0X" {
			valus = valus[:2]
		}
	}
	num, _ := new(big.Int).SetString(valus, 16)
	nums := num.Int64()
	return toAddress, float64(nums)
}

func BlogRefreshInit() error {
	hash := &dao.Blog{}
	if err := (hash).First(); err != nil {
		return err
	}
	blog := blogRefresh{
		Chan:   make(chan int),
		Number: hash.BlogNumber,
	}
	EthAddressInit()
	go (&blog).blockServer()
	go (&blog).hashServer()
	return nil
}
