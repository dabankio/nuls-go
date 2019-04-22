package nuls

type Inputs struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

type OutPuts struct {
	ToAddress string `json:"toAddress"`
	Amount    int64  `json:"amount"`
}
