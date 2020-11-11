package http

import (
	"fmt"
	"net/http"

	syslog "summitto.com/txsigner/shared/logging/system"
	"summitto.com/txsigner/shared/transport"
)

type (
	Transport struct {
		httpServer  *http.Server
		router      *http.ServeMux
		endpoints   []*Endpoint
		middlewares []*transport.Middleware
		logger      syslog.SystemLogger
	}

	TransportWithSysLog struct {
		transport *Transport
		logger    *syslog.SysLogger
	}
)

func NewTransport(
	endpoints []*Endpoint,
	middlewares []*transport.Middleware,
) *TransportWithSysLog {
	transport := &Transport{
		endpoints:   endpoints,
		middlewares: middlewares,
	}

	return &TransportWithSysLog{
		transport: transport,
		logger:    syslog.NewSysLogger(true),
	}
}

func (t *TransportWithSysLog) Init(port string) {
	t.transport.router = http.NewServeMux()
	t.transport.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: t.transport.router,
	}

	t.logger.Log("Initialsed http transport server and router on Addr:", t.transport.httpServer.Addr)
}

func (t *TransportWithSysLog) Register() {
	for _, endpoint := range t.transport.endpoints {
		t.transport.router.Handle(endpoint.url, *endpoint)
		t.logger.Log("Http transport registered endpoint:", endpoint.url)
	}
}

func (t *TransportWithSysLog) Start() (bool, error) {
	t.logger.Log("Starting the http transport server...")

	err := t.transport.httpServer.ListenAndServe()

	if err != nil {
		t.logger.Log("Failed to start the http transport server...")
		return false, err
	}

	t.logger.Log("Started the http transport server...")

	return true, nil
}

func (t *TransportWithSysLog) Stop() (bool, error) {
	t.logger.Log("Stopping the http transport server...")

	err := t.transport.httpServer.Close()

	if err != nil {
		t.logger.Log("Error while stopping the http transport server", err)
		return false, err
	}

	t.logger.Log("Successfuly stopped the http transport server", err)

	return true, nil
}
