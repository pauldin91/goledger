package models

import (
	"encoding/json"
	"time"

	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

type TransactionDto struct {
	TxID      string                `json:"txid"`
	Sender    string                `json:"sender"`
	Recipient string                `json:"recipient"`
	Amount    float64               `json:"amount"`
	Timestamp time.Time             `json:"timestamp"`
	Signature string                `json:"signature"`
	PublicKey string                `json:"pubkey"`
	TxInputs  map[string]tx.TxInput `json:"inputs"`
	TxOutputs []tx.TxOutput         `json:"outputs"`
}

func (transaction TransactionDto) IsValid() bool {
	outs, _ := json.Marshal(transaction.TxOutputs)
	var tsString string = utils.Hash(string(outs))
	return utils.VerifySignature(transaction.PublicKey, []byte(tsString), []byte(transaction.Signature))
}
