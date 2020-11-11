package transactions

type (
	HttpHandler interface {
		GetPubKey(request interface{}) (interface{}, error)
		PutTransaction(request interface{}) (interface{}, error)
		PostSignature(request interface{}) (interface{}, error)
	}

	http struct {
		processor Processor
	}

	getPubKeyRequest struct{}

	getPubKeyResponse struct {
		PublicKey string
	}

	putTransactionRequest struct {
		tx []byte
	}

	putTransactionResponse struct {
		ID string
	}

	postSignatureRequest struct {
		ids []string
	}

	postSignatureResponse struct {
		Message   []string
		Signature string
	}

	// look into adding transformers as functions that take in requets/responses and return them transformed
)

func (t *http) GetPubKey(request interface{}) (interface{}, error) {
	pubkey := t.processor.GetPublicKey()
	response := GetPubKeyResponse(pubkey)
	return response, nil
}

func (t *http) PutTransaction(request interface{}) (interface{}, error) {
	rqst := request.(map[string]interface{})
	r := PutTransactionRequest(rqst["tx"])
	id := t.processor.SaveTransaction(r.tx)
	response := PutTransactionResponse(id)
	return response, nil
}

func (t *http) PostSignature(request interface{}) (interface{}, error) {
	rqst := request.(map[string]interface{})
	r := PostSignatureRequest(rqst["ids"])
	txs, signature := t.processor.SignTransactions(r.ids)
	response := PostSignatureResponse(txs, signature)
	return response, nil
}

func NewHttpHandler() *http {
	return &http{
		processor: NewProcessor(),
	}
}
