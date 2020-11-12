package main

import (
	dcfg "transcryptoapi/config/daemon"
	txHttpService "transcryptoapi/protocols/http/transactions"
)

func main() {
	service := txHttpService.NewTxHttpService()

	service.Configure(dcfg.Config.Port)
	service.Run()
	// do we need a blocker for system in here?
	defer service.Stop()
}
