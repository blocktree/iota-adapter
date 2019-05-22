/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package iotaledger

import (
	"fmt"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/openwallet/openwallet"
	"github.com/iotaledger/iota.go/api"
)

const (
	maxAddresNum = 10000
)

// WalletManager is wallet manger
type WalletManager struct {
	openwallet.AssetsAdapterBase
	openw.WalletManager

	API             *api.API
	Config          *WalletConfig
	Log             *log.OWLogger
	Decoder         openwallet.AddressDecoder
	TxDecoder       openwallet.TransactionDecoder
	ContractDecoder openwallet.SmartContractDecoder
	Blockscanner    openwallet.BlockScanner
}

// NewWalletManager is init wallet function
func NewWalletManager() *WalletManager {
	wm := WalletManager{}

	wm.Config = NewConfig(Symbol)
	log := log.NewOWLogger(wm.Symbol())

	iotaAPI, err := api.ComposeAPI(api.HTTPClientSettings{URI: wm.Config.ServerAPI})
	if err != nil {
		// handle error
		log.Error("ComposeAPI error occurred unexpected.")
	}

	wm.API = iotaAPI
	wm.Log = log
	return &wm
}

// CreateAssetsAccount does not supported.
func (wm *WalletManager) CreateAssetsAccount(appID, walletID, password string, account *openwallet.AssetsAccount, otherOwnerKeys []string) (*openwallet.AssetsAccount, *openwallet.Address, error) {
	log.Error("not impl.")
	return nil, nil, fmt.Errorf("Not impl")
}

// GetAssetsAccountList does not supported.
func (wm *WalletManager) GetAssetsAccountList(appID, walletID string, offset, limit int) ([]*openwallet.AssetsAccount, error) {
	log.Error("not impl.")
	return nil, fmt.Errorf("Not impl")
}
