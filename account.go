package nuls

import (
	"fmt"
	"net/http"
)

const (
	accountPath = rootPath + "/account"
)

// GetAccount Get account information given an account ID.
func (c *Client) GetAccount(account string) (info AccountInfo, err error) {
	err = c.call(http.MethodGet, accountPath, "", nil, &info)
	return
}

// 创建账户
func (c *Client) CreateAccount(count int, password string) (info AccountList, err error) {
	param := M{
		"count":    count,
		"password": password,
	}
	err = c.call(http.MethodPost, accountPath, "", param, &info)
	return
}

// [余额] 查询本地所有账户总余额
func (c *Client) Balance() (info AccountBalance, err error) {
	option := "/balance"
	err = c.call(http.MethodGet, accountPath, option, nil, &info)
	return
}

// api/account/validate/{address} [验证地址格式是否正确]
func (c *Client) ValidateAddress(address string) (info ValidateAddress, err error) {
	option := fmt.Sprintf("/validate/%s", address)
	err = c.call(http.MethodGet, accountPath, option, nil, &info)
	return
}
