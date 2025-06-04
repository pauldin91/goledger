package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pauldin91/goledger/src/blockchain"
)

func TestAddTransactionToThePool(t *testing.T) {
	pool.AddOrUpdateById(*transaction)
	tt := pool.TransactionById(transaction.Id.String())
	if tt == nil {
		t.Error("transaction was not found in the pool")
		return
	}

	if tt.Input.Signature != transaction.Input.Signature {
		t.Error("invalid input in the transaction")
	}

	if tt.Output[senderWallet.Address].String() != transaction.Output[senderWallet.Address].String() {
		t.Error("Inputs dont much")
	}

	pool.Clear()

	if pool.Size() != 0 {
		t.Error("pool does not clear")
		return
	}
}

func TestValidTransactions(t *testing.T) {
	var t1 = blockchain.NewTransaction(senderWallet, recipientWallet.Address, 100)
	var t2 = *t1
	t2.Amount = -100
	t2.Id = uuid.New()

	pool.AddOrUpdateById(*t1)
	pool.AddOrUpdateById(t2)

	valid := pool.ValidTransactions()
	if len(valid) != 1 {
		t.Error("Invalid valid transactions length")
	}

}
