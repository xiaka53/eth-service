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
	gas *= math.Pow(10, 18)
	defer client.Close()
	var gasresult string
	if err = client.Call(&gasresult, "eth_gasPrice"); err != nil {
		return
	}
	gasresultBigInt, _ := new(big.Int).SetString(gasresult[2:], 16)
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
		Gas:   hexutil.EncodeUint64(uint64(gas) / gasresultBigInt.Uint64()),
		Gasp:  gasresult,
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

// //TODO 私钥转账法
//func (a *EthV1) Send(from, to string, value, gas float64) (hash string, err error) {
//	client := public.GetEthclient()
//	defer client.Close()
//	privateKey, err := crypto.HexToECDSA(from)
//	if err != nil {
//		return "", err
//	}
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		return "", err
//	}
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		return "", err
//	}
//	values := big.NewInt(0)
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		return "", err
//	}
//	toAddress := common.HexToAddress(to)
//	//tokenAddress := common.HexToAddress("contract address")
//	var data []byte
//	tx := types.NewTransaction(nonce, toAddress, values, uint64(gas)/(gasPrice.Uint64()), gasPrice, data)
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		return "", err
//	}
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		return "", err
//	}
//	err = client.SendTransaction(context.Background(), signedTx)
//	return
//}

func (a *EthV1) Transfer(address string, amont, from int) []public.Transfer {
	return (&Hash{Send: address}).Transfer(from-1, amont)
}

func (a *EthV1) EstimateGas(from, to string, value float64) float64 {
	client := public.GetEthclient()
	defer client.Close()
	value *= math.Pow(10, 18)
	tohax := common.HexToAddress(to)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0
	}
	msg := ethereum.CallMsg{
		From:     common.HexToAddress(from),
		To:       &tohax,
		Gas:      100000,
		GasPrice: gasPrice,
		Value:    big.NewInt(int64(value)),
		Data:     nil,
	}
	gas, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0
	}
	return float64(int64(gas)*gasPrice.Int64()) / math.Pow(10, 18)
}

//func (a *EthV1) EstimateGas(from, to string, value float64) float64 {
//	client := public.GetClient()
//	defer client.Close()
//	value *= math.Pow(10, 18)
//	trans := struct {
//		From  string `json:"from"`
//		To    string `json:"to"`
//		Value string `json:"value"`
//		Gas   string `json:"gas"`
//		Gasp  string `json:"gasp"`
//	}{
//		From:  from,
//		To:    to,
//		Value: hexutil.EncodeUint64(uint64(value)),
//		Gas:   hexutil.EncodeUint64(uint64(100000)),
//		Gasp:  hexutil.EncodeUint64(uint64(10000000000)),
//	}
//	var aaa interface{}
//	err := client.Call(&aaa, "eth_estimateGas", trans)
//	if err != nil {
//		return 0
//	}
//	return 1
//}

func (a *EthV1) HaxLog(hax string) public.Transfer {
	data := Hash{Hash: hax}
	return (&data).Hax("*")
}
