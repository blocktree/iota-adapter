package openwtester

import (
	"github.com/blocktree/iota-adapter/iotaledger"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(iotaledger.Symbol, iotaledger.NewWalletManager())
}
