package openwtester

import (
	"github.com/assetsadapterstore/ausd-adapter/ausd"
	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
)

func init() {
	//注册钱包管理工具
	log.Notice("Wallet Manager Load Successfully.")
	openw.RegAssets(ausd.Symbol, ausd.NewWalletManager())
}
