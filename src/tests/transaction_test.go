package tests

import (
	"testing"

	"github.com/pauldin91/goledger/src/utils"
)

func TestNewTransaction(t *testing.T) {

	if len(transaction.Output) != 2 {
		t.Error("Transaction output length is always 2")
	}

	sender, ex := transaction.Output[senderWallet.Address]
	if !ex ||
		sender.Amount != senderWallet.Balance-amount ||
		sender.Address != senderWallet.Address {
		t.Error("Invalid sender transaction")
	}

	recipient, ex := transaction.Output[recipientWallet.Address]
	if !ex ||
		recipient.Amount != recipientWallet.Balance+amount ||
		recipient.Address != recipientWallet.Address {
		t.Error("Invalid recipient transaction")
	}
}

func TestVerifyTransaction(t *testing.T) {
	res := transaction.Verify()
	if !res {
		t.Error("Valid transaction should be validated")
	}
	var copy utils.Input = transaction.Output[recipientWallet.Address]

	copy.Amount = 30000
	transaction.Output[recipientWallet.Address] = copy

	res = transaction.Verify()
	if res {
		t.Error("invalid transaction should be invalidated")
	}

}
