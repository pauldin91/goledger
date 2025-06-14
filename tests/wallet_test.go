package tests

import (
	"testing"

	"github.com/pauldin91/goledger/src/transaction"
	"github.com/pauldin91/goledger/src/utils"
)

func TestWalletCreation(t *testing.T) {
	var wallet = transaction.NewWallet()
	if wallet.GetAddress() != utils.Hash(wallet.GetPubKey()) {
		t.Errorf("address should be the hashed value of pubkey %s, intead was %s", utils.Hash(wallet.GetPubKey()), wallet.GetAddress())
	}
}

func TestBalance(t *testing.T) {
	senderWallet.WithUTXOs(utxoSet)
	balance := senderWallet.CalculateBalance()
	total := 0.0
	for _, v := range utxoSet {
		if v.Address == senderWallet.GetAddress() {
			total += v.Amount
		}
	}
	if total != balance {
		t.Errorf("Derived balance from utxos was %.10f while calculated balance %.10f\n", total, balance)
	}
}

func TestValidSendOutputToMempool(t *testing.T) {

	senderWallet.WithUTXOs(utxoSet)
	isSent := senderWallet.Send(validOutput, &tpool)

	if !isSent {
		t.Errorf("Not sent")
	}
}
