package transactions

type (
	Transaction struct {
		ID      string
		Content []byte
	}
)

func (t *Transaction) Valid() error {
	return nil
}
