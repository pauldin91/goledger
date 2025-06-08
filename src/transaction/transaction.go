package transaction

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id        uuid.UUID          `json:"id"`
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
		Id: uuid.New(),
	}
	return transaction
}

func (t *Transaction) sign() {
}

func NewTransaction(recipient string, amount float64) *Transaction {
	outputs := []TxInput{}
	var created Transaction = transactionWithOutputs(outputs, amount)
	return &created
}

func (t *Transaction) Update(recipientAddress string, amount float64) {

}

func (transaction *Transaction) Verify() bool {
	return true
}
