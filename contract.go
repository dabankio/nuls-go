package nuls

import (
	"fmt"
	"net/http"
)

const (
	contractPath = rootPath + "/contract"
)

// /api/contract/{address} 验证是否为合约地址
func (c *Client) ValidateContractAddress(address string) (info ValidateContract, err error) {
	option := fmt.Sprintf("/%s", address)
	err = c.call(http.MethodGet, contractPath, option, nil, &info)
	return
}

// /api/contract/transfer 向智能合约地址转账
func (c *Client) ContractTransfer(address, toAddress, password string, gasLimit, price int, amount int64) (info Transfer, err error) {
	param := M{
		"address":   address,
		"toAddress": toAddress,
		"gasLimit":  gasLimit,
		"price":     price,
		"password":  password,
		"amount":    amount,
	}
	option := "/transfer"
	err = c.call(http.MethodPost, contractPath, option, param, &info)
	return
}
