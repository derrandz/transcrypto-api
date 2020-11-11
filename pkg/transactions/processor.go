package transactions

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"log"

	dcfg "summitto.com/txsigner/config/daemon"
)

type (
	Processor interface {
		GetPublicKey() string
		SaveTransaction(data []byte) string
		SignTransactions(ids []string) ([]string, string)
	}

	processor struct {
		transactions TxMemoryStore
	}
)

func (p *processor) GetPublicKey() string {
	publicKeyBase64, _, err := GetDaemonPublicKey(dcfg.Config.PrivateKey)
	if err != nil {
		log.Fatal("Error while trying to retrieve the public key")
		return ""
	}
	return publicKeyBase64
}

func (p *processor) SaveTransaction(data []byte) string {
	return p.transactions.SaveTransaction(data)
}

func (p *processor) SignTransactions(ids []string) ([]string, string) { // fix me
	txs := p.transactions.GetTransactionsById(ids)
	txsArr := make([]string, len(txs))

	for i, tx := range txs {
		txsArr[i] = base64.StdEncoding.EncodeToString(tx)
	}

	writer := bytes.NewBufferString("")
	json.NewEncoder(writer).Encode(txsArr)

	_, privateKey, err := GetPrivateKey(dcfg.Config.PrivateKey)
	if err != nil {
		log.Fatal("Error while trying to decode the private key")
		return make([]string, 0), ""
	}

	message := []byte(writer.String())

	signatureData := ed25519.Sign(privateKey, message)
	signature := base64.StdEncoding.EncodeToString(signatureData)

	return txsArr, signature
}

func NewProcessor() Processor {
	return &processor{
		transactions: NewTransactionStore(),
	}
}
