package transaction

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/pauldin91/goledger/src/models"
	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

type Transaction struct {
	TxID      string
	Sender    string
	Recipient string
	Amount    float64
	Timestamp time.Time
	Signature string
	PublicKey string
	TxInputs  map[string]tx.TxInput
	TxOutputs []tx.TxOutput
}

func (t Transaction) String() string {
	jsonT, _ := json.Marshal(t)
	return string(jsonT)
}
func transactionWithOutputs(inputs []tx.TxInput, output tx.TxOutput) Transaction {
	transaction := Transaction{
		Amount: output.Amount,
	}
	return transaction
}

func (t *Transaction) Sign() {
}

func NewTransaction(output tx.TxOutput) *Transaction {
	var created Transaction = transactionWithOutputs([]tx.TxInput{}, output)
	return &created
}

func (ts Transaction) IsValid() bool {
	outs, _ := json.Marshal(ts.TxOutputs)
	var tsString string = utils.Hash(string(outs))
	return utils.VerifySignature(ts.PublicKey, []byte(tsString), []byte(ts.Signature))
}

func (transaction Transaction) Hash() string {

	total := strconv.FormatFloat(transaction.Amount, 'f', -1, 64) + transaction.Recipient + transaction.Sender + transaction.Timestamp.String()
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
		Sender:    transaction.Sender,
		Recipient: transaction.Recipient,
		Amount:    transaction.Amount,
		Signature: transaction.Signature,
		Timestamp: transaction.Timestamp,
		PublicKey: transaction.PublicKey,
		TxInputs:  transaction.TxInputs,
		TxOutputs: transaction.TxOutputs,
	}
}
