package TransactionsHttpService

import (
	tx "summitto.com/txsigner/pkg/transactions"
	"summitto.com/txsigner/shared/transport/http"
)

type (
	txHttpService struct {
		transport *http.TransportWithSysLog
	}
)

func NewTxHttpService() *txHttpService {
	httpHandler := tx.NewHttpHandler()

	endpoints := []*http.Endpoint{
		http.NewJsonHttpEndpoint(httpHandler.GetPubKey, "/public_key", "GET"),
		http.NewJsonHttpEndpoint(httpHandler.PutTransaction, "/transaction", "PUT"),
		http.NewJsonHttpEndpoint(httpHandler.PostSignature, "/signature", "POST"),
	}

	httpTransport := http.NewTransport(endpoints, nil)

	return &txHttpService{
		transport: httpTransport,
	}
}

func (s *txHttpService) Configure(port string) {
	s.transport.Init(port)
	s.transport.Register()
}

func (s *txHttpService) Run() (bool, error) {
	s.transport.Start()
	return true, nil
}

func (s *txHttpService) Stop() (bool, error) {
	return s.transport.Stop()
}
