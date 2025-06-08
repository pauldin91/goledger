package transaction

type TransactionPool struct {
	Transactions map[string]Transaction
}

func (p *TransactionPool) Size() int {
	return len(p.Transactions)
}

func (p *TransactionPool) AddOrUpdateById(transaction Transaction) {
	if transaction.Amount > 0 {
		p.Transactions[transaction.TxID] = transaction
	}

}

func (p *TransactionPool) TransactionById(id string) *Transaction {
	tr, ok := p.Transactions[id]
	if ok {
		return &tr
	}
	return nil
}

func (p *TransactionPool) ValidTransactions() []Transaction {
	var transactions []Transaction
	for _, t := range p.Transactions {
		transaction := filter(t)
		if transaction != nil {
			transactions = append(transactions, *transaction)
		}
	}
	return transactions
}
func (p *TransactionPool) Clear() {
	p.Transactions = make(map[string]Transaction)
}
