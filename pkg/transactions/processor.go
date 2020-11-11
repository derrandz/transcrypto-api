package transactions

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
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
	publicKeyBase64, _, err := GetDaemonPublicKey()
	if err != nil {
		log.Fatal("Error while trying to retrieve the public key")
		return ""
	}
	return publicKeyBase64
}

func (p *processor) SaveTransaction(data []byte) string {
	return p.transactions.SaveTransaction(data)
}

func (p *processor) SignTransactions(ids []string) ([]string, string) {
	txs := p.transactions.GetTransactionsById(ids)
	txsArr := make([]string, len(txs))

	for _, tx := range txs {
		txsArr = append(txsArr, base64.StdEncoding.EncodeToString(tx))
	}

	fmt.Printf("txsArr: %v, txs: %v", txsArr, txs)

	writer := bytes.NewBufferString("")
	json.NewEncoder(writer).Encode(txsArr)

	_, privateKey, err := GetDaemonPrivateKey()
	if err != nil {
		log.Fatal("Error while trying to decode the private key")
		return make([]string, 0), ""
	}

	message := []byte(writer.String())

	fmt.Printf("Signed json transactions as message: %s", message)

	signature := ed25519.Sign(privateKey, message)

	fmt.Printf("Signature: %s", message)

	return txsArr, string(signature)
}

func NewProcessor() Processor {
	return &processor{
		transactions: NewTransactionStore(),
	}
}
