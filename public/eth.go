package public

type Eth interface {
	GetBalance(string) (float64, error)
	Send(string, string, float64, float64) (Transfer, error)
	Transfer(string, int, int) []Transfer
	EstimateGas(string, string, float64) float64
	HaxLog(string) Transfer
}

type Transfer struct {
	BlokNum    int     `json:"blok_num"`
	Hash       string  `json:"hash"`
	Send       string  `json:"send"`
	To         string  `json:"to"`
	Token      string  `json:"token"`
	TokenName  string  `json:"token_name"`
	Num        float64 `json:"num"`
	Gas        string  `json:"gas"`
	GasPrice   string  `json:"gas_price"`
	Createtime int     `json:"createtime"`
}
