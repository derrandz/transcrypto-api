package transactions

import (
	"fmt"
)

func GetPubKeyResponse(pubkey string) *getPubKeyResponse {
	return &getPubKeyResponse{
		PublicKey: pubkey,
	}
}

func PutTransactionRequest(data interface{}) *putTransactionRequest {
	r := new(putTransactionRequest)
	rData := data.([]interface{})

	txData := []byte(rData[0].(string))

	r.tx = txData

	return r
}

func PutTransactionResponse(id string) *putTransactionResponse {
	return &putTransactionResponse{
		ID: id,
	}
}

func PostSignatureRequest(reqIDs interface{}) *postSignatureRequest {
	r := new(postSignatureRequest)
	IDs := reqIDs.([]interface{})
	ids := make([]string, len(IDs))

	for _, v := range IDs {
		id := fmt.Sprintf("%v", v)
		ids = append(ids, id)
	}

	r.ids = ids

	return r
}

func PostSignatureResponse(txs []string, signature string) *postSignatureResponse {
	return &postSignatureResponse{message: txs, signature: signature}
}
