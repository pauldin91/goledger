package tests

import (
	"encoding/json"
	"testing"
)

func TestCreateTransaction(t *testing.T) {

	executed := senderWallet.CreateTransaction(recipientWallet.Address, testAmounts[0].amount, *bc, &tp)
	if len(tp.Transactions) != 1 || !executed {
		t.Errorf("should have %d but have %d\n", 1, len(tp.Transactions))
	}
	tp.Clear()
}

func TestBalance(t *testing.T) {

	var senderBalance float64 = senderWallet.Balance
	var recipientBalance float64 = recipientWallet.Balance
	for _, ta := range testAmounts {
		executed := senderWallet.CreateTransaction(recipientWallet.Address, ta.amount, *bc, &tp)
		if executed != ta.shouldBeExecuted {
			t.Errorf("test with amount %0.8f it was supposed to %v while got %v", ta.amount, ta.shouldBeExecuted, executed)
			continue
		} else if !executed {
			continue
		}

		jsonTransactions, _ := json.Marshal(tp.Transactions)

		bc.AddBlock(string(jsonTransactions))

		senderBalance = senderWallet.Balance
		senderWallet.Balance = senderWallet.CalculateBalance(*bc)
		if senderWallet.Balance != senderBalance-ta.amount {
			t.Errorf("Sender should have a balance of %0.8f but has %0.8f\n", senderBalance-ta.amount, senderWallet.Balance)
		}
		recipientBalance = recipientWallet.Balance
		recipientWallet.Balance = recipientWallet.CalculateBalance(*bc)
		if recipientWallet.Balance != recipientBalance+ta.amount {
			t.Errorf("Recipient should have a balance of %0.8f but has %0.8f\n", recipientBalance+ta.amount, recipientWallet.Balance)
		}

	}
	tp.Clear()
}
