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
		return
	}

	response, err := e.funcHandler(request)
	if err != nil {
		fmt.Println("Error in funchandler", err)
		return
	}

	if err := e.encode(rw, response); err != nil {
		fmt.Println("Error in encoding", err)
		return
	}
}

func JsonResponseEncoder(w http.ResponseWriter, response interface{}) error {
	fmt.Println(response)
	return json.NewEncoder(w).Encode(response)
}

func JsonRequestDecoder(r *http.Request) (response interface{}, err error) {
	var request interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
