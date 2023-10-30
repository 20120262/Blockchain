package Blockchain

type Transaction struct {
	Data []byte
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{data}
}
