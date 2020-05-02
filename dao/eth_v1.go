package dao

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/xiaka53/eth-service/public"
	"math"
	"math/big"
)

type EthV1 struct {
}

func NewEth() *EthV1 {
	return &EthV1{}
}

func (a *EthV1) GetBalance(address string) (value float64, err error) {
	var (
		balanceAt *big.Int
	)
	client := public.GetEthclient()
	defer client.Close()
	account := common.HexToAddress(address)
	if balanceAt, err = client.BalanceAt(context.Background(), account, nil); err != nil {
		return
	}
	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	value, _ = ethValue.Float64()
	return
}

func (a *EthV1) Send(from, to string, value, gas float64) (hash string, err error) {
	value *= math.Pow(10, 18)
	client := public.GetClient()
	defer client.Close()
	trans := struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Value string `json:"value"`
		Gas   string `json:"gas"`
		Gasp  string `json:"gasp"`
	}{
		From:  from,
		To:    to,
		Value: hexutil.EncodeUint64(uint64(value)),
		Gas:   hexutil.EncodeUint64(uint64(100000)),
		Gasp:  hexutil.EncodeUint64(uint64(10000000000)),
	}
	var result bool
	if err = client.Call(&result, "personal_unlockAccount", from, "Z*DHJ%IOlGtJh5TFng3pt3mRD^Q9II!&sCDDpzT3vAFRROVPp$BMzO$1Bf4P6GEF"); err != nil {
		return
	}
	err = client.Call(&hash, "eth_sendTransaction", trans)
	_ = client.Call(&result, "personal_lockAccount", from)
	if err != nil {
		if err.Error() == "insufficient funds for gas * price + value" {
			err = errors.New("余额不足")
		}
	}
	return
}

func (a *EthV1) Transfer(address string, amont, from int) []public.Transfer {
	return (&Hash{Send: address}).Transfer(from-1, amont)
}

func (a *EthV1) EstimateGas(from, to string, value float64) float64 {
	client := public.GetEthclient()
	defer client.Close()
	value *= math.Pow(10, 18)
	tohax := common.HexToAddress(to)
	msg := ethereum.CallMsg{
		From:     common.HexToAddress(from),
		To:       &tohax,
		Gas:      100000,
		GasPrice: big.NewInt(10000000000),
		Value:    big.NewInt(int64(value)),
		Data:     nil,
	}
	gas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0
	}
	return float64(gas) / math.Pow(10, 18)
}

func (a *EthV1) HaxLog(hax string) public.Transfer {
	data := Hash{Hash: hax}
	return (&data).Hax("*")
}
