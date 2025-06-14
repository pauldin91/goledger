package models

import (
	"time"

	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

type TransactionDto struct {
	TxID      string        `json:"txid"`
	Sender    string        `json:"sender"`
	Recipient string        `json:"recipient"`
	Amount    float64       `json:"amount"`
	Timestamp time.Time     `json:"timestamp"`
	Signature string        `json:"signature"`
	PublicKey string        `json:"pubkey"`
	TxInputs  []tx.TxInput  `json:"inputs"`
	TxOutputs []tx.TxOutput `json:"outputs"`
}

func (transaction TransactionDto) Hash() string {

	total := transaction.Timestamp.String()

	inputs := ""
	for _, v := range transaction.TxInputs {
		inputs += v.Hash()
	}
	outputs := ""
	for _, v := range transaction.TxOutputs {
		outputs += v.Hash()
	}
	total += inputs + outputs
	return utils.Hash(total)
}

func (transaction TransactionDto) IsValid() bool {
	var tsString string = transaction.Hash()
	return utils.VerifySignature(transaction.PublicKey, []byte(tsString), []byte(transaction.Signature))
}
