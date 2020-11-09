package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"summitto.com/txsigner/shared/transport"
)

type (
	Endpoint struct {
		funcHandler transport.EndpointHandler
		encode      transport.EncoderFunc
		decode      transport.DecodeFunc
		url         string
		method      string
	}
)

func NewHttpEndpoint(
	endpoint transport.EndpointHandler,
	encoder transport.EncoderFunc,
	decoder transport.DecodeFunc,
	url string,
	method string,
) *Endpoint {
	httpEndpoint := &Endpoint{
		funcHandler: endpoint,
		encode:      encoder,
		decode:      decoder,
		url:         url,
		method:      method,
	}

	return httpEndpoint
}

func NewJsonHttpEndpoint(
	endpoint transport.EndpointHandler,
	url string,
	method string,
) *Endpoint {
	httpEndpoint := &Endpoint{
		funcHandler: endpoint,
		encode:      JsonResponseEncoder,
		decode:      JsonRequestDecoder,
		url:         url,
		method:      method,
	}

	return httpEndpoint
}

func (e Endpoint) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method != e.method {
		fmt.Println("Not the right method though")
		rw.Write([]byte("404"))
		return
	}

	request, err := e.decode(r)
	if err != nil {
		fmt.Println("Error in decoding", err)
		// I entertain the idea of having an error handler for each request (useful)
		// s.errorHandler.Handle(ctx, err)
		// s.errorEncoder(ctx, err, w)
		return
	}

	response, err := e.funcHandler(request)
	if err != nil {
		fmt.Println("Error in funchandler", err)
		// same same for error handler
		// s.errorHandler.Handle(ctx, err)
		// s.errorEncoder(ctx, err, w)
		return
	}

	if err := e.encode(rw, response); err != nil {
		fmt.Println("Error in encoding", err)
		// s.errorHandler.Handle(ctx, err)
		// s.errorEncoder(ctx, err, w)
		return
	}

	rw.Write([]byte("Hello world"))
}

func JsonResponseEncoder(w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func JsonRequestDecoder(r *http.Request) (response interface{}, err error) {
	var request interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
