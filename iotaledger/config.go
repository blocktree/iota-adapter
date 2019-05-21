/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package iotaledger

import (
	"path/filepath"
	"strings"

	"github.com/blocktree/openwallet/common/file"
)

const (
	Symbol        = "IOTA"
	defaultConfig = `

# Node server api
serverAPI = "http://127.0.0.1:14265"

`
)

// WalletConfig is the wallet basic config
type WalletConfig struct {

	//币种
	Symbol string
	//配置文件路径
	configFilePath string
	//配置文件名
	configFileName string
	//账本数据文件
	ledgerFile string
	//本地数据库文件路径
	dbPath string
	//默认配置内容
	DefaultConfig string
	//node server
	ServerAPI string
}

// NewConfig is wallet init method
func NewConfig(symbol string) *WalletConfig {

	c := WalletConfig{}

	//币种
	c.Symbol = symbol
	c.configFilePath = filepath.Join("conf")
	c.configFileName = c.Symbol + ".ini"
	c.ledgerFile = "ledger.db"
	c.dbPath = filepath.Join("data", strings.ToLower(c.Symbol), "db")
	c.ServerAPI = "http://127.0.0.1:14265"
	c.DefaultConfig = defaultConfig

	//创建目录
	file.MkdirAll(c.dbPath)

	return &c
}
