package transactions

import (
	"fmt"

	"github.com/google/uuid"
	"summitto.com/txsigner/shared/store"
	"summitto.com/txsigner/shared/store/memory"
)

type (
	TxMemoryStore interface {
		GetTransactionById(id string) []byte
		GetTransactionsById(ids []string) [][]byte
		SaveTransaction(data []byte) string
	}

	txMemoryStore struct {
		mem store.Memory
	}
)

func (ts *txMemoryStore) GetTransactionById(id string) []byte {
	txdata, correct := (ts.mem.Get(id)).([]byte)
	if correct { // implement error on incorrect type
		return txdata
	}

	return make([]byte, 0)
}

func (ts *txMemoryStore) GetTransactionsById(ids []string) [][]byte {
	txs := make([][]byte, len(ids))
	for i, id := range ids {
		tx, _ := (ts.mem.Get(id)).([]byte)
		txs[i] = []byte(tx)
	}
	return txs
}

func (ts *txMemoryStore) SaveTransaction(data []byte) string {
	randomID := fmt.Sprintf("%v", uuid.New())
	created := ts.mem.Set(randomID, data)
	if created {
		return randomID
	}
	return ""
}

func NewTransactionStore() TxMemoryStore {
	return &txMemoryStore{
		mem: memory.NewMemoryStore(),
	}
}
