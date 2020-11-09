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
		data []byte
	}

	putTransactionResponse struct {
		id string
	}

	postSignatureRequest struct {
		ids []string
	}

	postSignatureResponse struct {
		message   [][]byte
		signature string
	}

	// look into adding transformers as functions that take in requets/responses and return them transformed
)

func (t *http) GetPubKey(request interface{}) (interface{}, error) {
	pubkey := t.processor.GetPublicKey()
	return getPubKeyResponse{PublicKey: pubkey}, nil
}

func (t *http) PutTransaction(request interface{}) (interface{}, error) {
	r := putTransactionRequest{data: request.([]byte)}
	id := t.processor.SaveTransaction(r.data)
	return putTransactionResponse{id: id}, nil
}

func (t *http) PostSignature(request interface{}) (interface{}, error) {
	r := postSignatureRequest{ids: request.([]string)}
	txs, signature := t.processor.SignTransactions(r.ids)
	return postSignatureResponse{message: txs, signature: signature}, nil
}

func NewHttpHandler() *http {
	return &http{
		processor: NewProcessor(),
	}
}
