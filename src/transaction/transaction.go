package transaction

import (
	"encoding/json"
	"time"
)

type Transaction struct {
	TxID      string             `json:"txid"`
	Sender    string             `json:"sender"`
	Recipient string             `json:"recipient"`
	Amount    float64            `json:"amount"`
	Timestamp time.Time          `json:"timestamp"`
	Signature string             `json:"signature"`
	PublicKey string             `json:"pubkey"`
	TxInputs  map[string]TxInput `json:"inputs"`
	TxOutputs []TxOutput         `json:"outputs"`
}

func (t Transaction) String() string {
	jsonT, _ := json.Marshal(t)
	return string(jsonT)
}
func transactionWithOutputs(outputs []TxInput, amount float64) Transaction {
	transaction := Transaction{
		Amount: amount,
	}
	return transaction
}

func (t *Transaction) sign() {
}

func NewTransaction(recipient string, amount float64) *Transaction {
	var created Transaction = transactionWithOutputs([]TxInput{}, amount)
	return &created
}

func (t *Transaction) Update(recipientAddress string, amount float64) {

}

func (transaction *Transaction) Verify() bool {
	return true
}
