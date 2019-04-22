package nuls

import (
	"fmt"
	"net/http"
)

const (
	transactionPath = rootPath + "/tx"
)

// /api/tx/hash/{hash} 根据hash查询交易
func (c *Client) GetTransaction(txId string) (info Transaction, err error) {
	option := fmt.Sprintf("/hash/%s", txId)
	err = c.call(http.MethodGet, transactionPath, option, nil, &info)
	return
}
