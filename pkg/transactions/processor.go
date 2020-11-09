package transactions

type (
	Processor interface {
		GetPublicKey() string
		SaveTransaction(data []byte) string
		SignTransactions(ids []string) ([][]byte, string)
	}

	processor struct {
		transactions TxMemoryStore
	}
)

func (p *processor) GetPublicKey() string {
	return ""
}

func (p *processor) SaveTransaction(data []byte) string {
	return p.transactions.SaveTransaction(data)
}

func (p *processor) SignTransactions(ids []string) ([][]byte, string) {
	txs := p.transactions.GetTransactionsById(ids)
	signature := ""

	//signing logic

	return txs, signature
}

func NewProcessor() Processor {
	return &processor{
		transactions: NewTransactionStore(),
	}
}
