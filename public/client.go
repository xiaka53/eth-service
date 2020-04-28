package public

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/xiaka53/DeployAndLog/lib"
)

func GetClient() *rpc.Client {
	client, err := rpc.Dial(fmt.Sprintf("http://%s", lib.ConfBase.Base.WebUrl))
	if err != nil {
		return nil
	}
	return client
}

func GetEthclient() *ethclient.Client {
	client, err := ethclient.Dial(fmt.Sprintf("http://%s", lib.ConfBase.Base.WebUrl))
	if err != nil {
		return nil
	}
	return client
}
