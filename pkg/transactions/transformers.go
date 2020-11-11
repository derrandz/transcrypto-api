package transactions

import (
	"encoding/base64"
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

	txData := rData[0].(string)

	r.tx, _ = base64.StdEncoding.DecodeString(txData)

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

	for i, v := range IDs {
		ids[i] = fmt.Sprintf("%v", v)
	}

	r.ids = ids

	return r
}

func PostSignatureResponse(txs []string, signature string) *postSignatureResponse {
	return &postSignatureResponse{Message: txs, Signature: signature}
}
