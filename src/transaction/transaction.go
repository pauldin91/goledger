package transaction

import (
	"encoding/json"
	"time"

	"github.com/pauldin91/goledger/src/models"
	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

type Transaction struct {
	TxID      string
	Timestamp time.Time
	Signature string
	PublicKey string
	TxInputs  []tx.TxInput
	TxOutputs []tx.TxOutput
}

func (t Transaction) String() string {
	jsonT, _ := json.Marshal(t)
	return string(jsonT)
}

func (ts *Transaction) Sign(keyPair utils.KeyPair) {
	hashed := ts.Hash()
	ts.Signature = keyPair.Sign(hashed)
}

func CreateTransaction(pubkey string, recipients []tx.TxOutput, utxos []tx.UTXO) *Transaction {

	inputs := make([]tx.TxInput, len(utxos))
	for i, utxo := range utxos {
		inputs[i] = utxo.Map()
	}

	tx := &Transaction{
		TxInputs:  inputs,
		TxOutputs: recipients,
		PublicKey: pubkey,
	}
	tx.TxID = tx.Hash()

	return tx
}

func (transaction Transaction) Hash() string {

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
	return total
}

func (transaction Transaction) Map() models.TransactionDto {
	return models.TransactionDto{
		TxID:      transaction.Hash(),
		Signature: transaction.Signature,
		Timestamp: transaction.Timestamp,
		PublicKey: transaction.PublicKey,
		TxInputs:  transaction.TxInputs,
		TxOutputs: transaction.TxOutputs,
	}
}

func (ts Transaction) IsValid() bool {
	var tsString string = ts.Hash()
	return utils.VerifySignature(ts.PublicKey, []byte(tsString), []byte(ts.Signature))
}
