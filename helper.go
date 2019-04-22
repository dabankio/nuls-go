/*
 * Copyright (c) 2019. Dabank Authors
 * All rights reserved.
 */
package nuls

import (
	"math/big"
	"reflect"
)

const (
	// coinbase 交易
	TxTypeCoinbase = iota + 1
	// 转账交易
	TxTypeTransfer
	// 设置别名
	TxTypeAlias
	// 创建共识节点交易
	TxTypeRegisterAgent
	// 委托交易(加入共识)
	TxTypeJoinConsensus
	// 取消委托交易(退出共识)
	TxTypeCancelDeposit
	// 黄牌惩罚
	TxTypeYellowPunish
	// 红牌惩罚
	TxTypeRedPunish
	// 停止节点(删除共识节点)
	TxTypeStopAgent
	// 跨链转账交易
	TxTypeCrossChainTransfer
	// 注册链交易
	TxTypeRegisterChainAndAsset
	// 销毁链
	TxTypeDestroyChainAndAsset
	// 为链新增一种资产
	TxTypeAddAssetToChain
	// 删除链上资产
	TxTypeRemoveAssetFromChain
)

const (
	// 创建智能合约交易
	TxTypeCreateContract = iota + 100
	// 调用智能合约交易
	TxTypeCallContract
	// 删除智能合约交易
	TxTypeDeleteContract
)

// compose adds input to param, whose key is tag
// if input is nil, compose is a no-op.
func compose(param map[string]interface{}, tag string, input interface{}) {
	// simple situation
	if input == nil {
		return
	}

	// needs dig further
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Interface:
		if v.IsNil() {
			return
		}
	}

	param[tag] = input
}

// M is a type shorthand for param input
type M map[string]interface{}

// AmountToInt64 converts a raw atomic NULS balance to a more human readable format.
func AmountToInt64(xmr float64) int64 {
	x := big.NewFloat(xmr)
	x1, _ := x.Mul(x, big.NewFloat(1e8)).Int64()
	return x1
}

// AmountToFloat64 converts raw atomic NULS to a float64
func AmountToFloat64(xmr int64) float64 {
	return float64(xmr) / 1e8
}
