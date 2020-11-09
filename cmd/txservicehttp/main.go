package main

import (
	dcfg "summitto.com/txsigner/config/daemon"
	txHttpService "summitto.com/txsigner/protocols/http/transactions"
)

func main() {
	service := txHttpService.NewTxHttpService()

	service.Configure(dcfg.Config.Port)
	service.Run()
	// do we need a blocker for system in here?
	defer service.Stop()
}
