package task

import (
	"github.com/xiaka53/eth-service/dao"
	"github.com/xiaka53/eth-service/public"
	"sync"
	"time"
)

type EthAddressStruct struct {
	Address map[string]bool
	Mux     sync.RWMutex
}

type newEthAddressStruct struct {
	Chan  chan string
	Count int
	Mux   sync.RWMutex
}

var (
	EthAddress    EthAddressStruct
	newEthAddress newEthAddressStruct
)

func EthAddressInit() {
	EthAddress = EthAddressStruct{
		Address: (&dao.Address{Status: "Y"}).Find(),
		Mux:     sync.RWMutex{},
	}
	go (&newEthAddress).address()
}

func (n *newEthAddressStruct) address() {
	address := (&dao.Address{Status: "N"}).Find()
	newEthAddress = newEthAddressStruct{Chan: make(chan string), Count: 0}
	for k, _ := range address {
		n.set(k)
	}
	time.Sleep(2 * time.Second) //TODO 延迟两秒避免上边go携程未执行完
	if newEthAddress.Count < 20 {
		n.newAddress(20 - len(newEthAddress.Chan))
	}
}

func (n *newEthAddressStruct) newAddress(count int) {
	client := public.GetClient()
	defer client.Close()
	for i := count; i > 0; i-- {
		var address string
		if err := client.Call(&address, "personal_newAccount", "Z*DHJ%IOlGtJh5TFng3pt3mRD^Q9II!&sCDDpzT3vAFRROVPp$BMzO$1Bf4P6GEF"); err == nil {
			(&dao.Address{Address: address, Status: "N"}).Create()
			n.set(address)
		}
	}
}

func (n *newEthAddressStruct) set(address string) {
	go func() {
		newEthAddress.Mux.Lock()
		newEthAddress.Count++
		newEthAddress.Mux.Unlock()
		newEthAddress.Chan <- address
	}()
}

func GetNewAddress() (address string) {
	var count int
	newEthAddress.Mux.Lock()
	newEthAddress.Count--
	count = newEthAddress.Count
	newEthAddress.Mux.Unlock()
	address = <-newEthAddress.Chan
	if err := (&dao.Address{Status: "N", Address: address}).Update(); err != nil {
		return ""
	}
	EthAddress.Mux.Lock()
	EthAddress.Address[address] = true
	EthAddress.Mux.Unlock()
	if count <= 10 {
		go (&newEthAddress).newAddress(20 - count)
	}
	return
}
