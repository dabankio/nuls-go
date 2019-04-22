package nuls

import "net/http"

const (
	accountLedgerPath = rootPath + "/accountledger"
)

// /api/accountledger/transfer 转账
func (c *Client) Transfer(address, toAddress, password string, amount int64) (info Transfer, err error) {
	param := M{
		"address":   address,
		"toAddress": toAddress,
		"password":  password,
		"amount":    amount,
	}
	option := "/transfer"
	err = c.call(http.MethodPost, accountLedgerPath, option, param, &info)
	return
}
