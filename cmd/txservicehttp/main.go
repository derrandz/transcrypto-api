package main

import (
	dcfg "summitto.com/txsigner/config/daemon"
	txHttpService "summitto.com/txsigner/protocols/http/transactions"
)

func main() {
	config := dcfg.GetDaemonConfig()
	service := txHttpService.NewTxHttpService()

	service.Configure(config.Port)
	service.Run()
	// do we need a blocker for system in here?
	defer service.Stop()
}
