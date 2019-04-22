package nuls

import (
	"fmt"
	"net/http"
)

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

// /api/accountledger/multipleAddressTransfer 多地址转账
func (c *Client) MultipleAddressTransfer(inputs []Inputs, outPuts []OutPuts) (info Transfer, err error) {
	param := M{
		"inputs":  inputs,
		"outputs": outPuts,
	}
	option := "/multipleAddressTransfer"
	err = c.call(http.MethodPost, accountLedgerPath, option, param, &info)
	return
}

// /api/accountledger/balance/{address} 账户地址查询账户余额
func (c *Client) AddressBalance(address string) (info Balance, err error) {
	option := fmt.Sprintf("/balance/%s", address)
	err = c.call(http.MethodGet, accountLedgerPath, option, nil, &info)
	return
}
