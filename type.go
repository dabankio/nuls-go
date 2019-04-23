package nuls

type Input struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

type OutPut struct {
	ToAddress string `json:"toAddress"`
	Amount    int64  `json:"amount"`
}
