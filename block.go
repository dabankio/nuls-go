package nuls

import "net/http"

const (
	blockPath = rootPath + "/block"
)

// 根据区块高度查询区块列表，包含区块打包的所有交易信息
func (c *Client) Block(startHeight, size int) (info BlockInfo, err error) {
	param := M{
		"startHeight": startHeight,
		"size":        size,
	}
	err = c.call(http.MethodGet, blockPath, "", param, &info)
	return
}

// "/api/block/newest/height" 查询最新区块高度
func (c *Client) BlockHeight() (info BlockHeight, err error) {
	option := "/newest/height"
	err = c.call(http.MethodGet, blockPath, option, nil, &info)
	return
}
