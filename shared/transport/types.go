package transport

import "net/http"

type (
	Transport interface {
		Init()
		Register()
		Start() (bool, error)
		Stop() (bool, error)
	}

	EndpointHandler func(request interface{}) (response interface{}, err error)
	DecodeFunc      func(request *http.Request) (response interface{}, err error)
	EncoderFunc     func(serverResponse http.ResponseWriter, response interface{}) (err error)
	Middleware      func(EndpointHandler) EndpointHandler
)
