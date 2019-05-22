package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/blocktree/iota-adapter/iotaledger"

	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

var (
	testApp        = "iota-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *iotaledger.WalletManager {
	log.SetLogFuncCall(true)
	// tc := openw.NewConfig()

	// tc.ConfigDir = configFilePath
	// tc.EnableBlockScan = false
	// tc.SupportAssets = []string{
	// 	"IOTA",
	// }
	return iotaledger.NewWalletManager()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO IOTA", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("wallet:", nw)
	t.Log("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "WHQF3H2Hqa2Pksp8vWmBDZpS7piEGVivRp")
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	t.Log("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		t.Log("wallet[", i, "] :", w)
	}
	t.Log("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "VzhK2fEviqi3UaG7h852Wdpw4kUsZYQC1K"
	account := &openwallet.AssetsAccount{Alias: "mainnetIOTA", WalletID: walletID, Required: 1, Symbol: "IOTA", IsTrust: true}
	_, _, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err == nil {
		t.Error(err)
		return
	}

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WHQF3H2Hqa2Pksp8vWmBDZpS7piEGVivRp"
	_, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err == nil {
		t.Error(err)
		return
	}

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WHQF3H2Hqa2Pksp8vWmBDZpS7piEGVivRp"
	accountID := "HgRBsaiKgoVDagwezos496vqKQCh41pY44JbhW65YA8t"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 5)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WHQF3H2Hqa2Pksp8vWmBDZpS7piEGVivRp"
	accountID := "HgRBsaiKgoVDagwezos496vqKQCh41pY44JbhW65YA8t"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		t.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		t.Log("address[", i, "] :", w.Address)
	}
	t.Log("address count:", len(list))

	tm.CloseDB(testApp)
}
