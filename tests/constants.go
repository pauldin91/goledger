package tests

import (
	"github.com/pauldin91/goledger/src/block"
	"github.com/pauldin91/goledger/src/pool"
	"github.com/pauldin91/goledger/src/transaction"
	"github.com/pauldin91/goledger/src/tx"
)

var genesisBlock = block.Genesis()

var tp = pool.MemPool{}

var amount float64 = 10.0

var bc = block.Create()

var senderWallet = transaction.NewWallet()
var recipientWallet = transaction.NewWallet()

var testAmounts = []struct {
	amount           float64
	shouldBeExecuted bool
}{
	{5.0, true},
	{11.0, true},
	{22.0, true},
	{-19.0, false},
	{50000.0, false},
}

var utxoSet []tx.UTXO = []tx.UTXO{
	{TxID: "tx1", OutputIndex: 0, Amount: 2.0, Address: senderWallet.GetAddress()},
	{TxID: "tx2", OutputIndex: 1, Amount: 1.0, Address: senderWallet.GetAddress()},
	{TxID: "tx3", OutputIndex: 2, Amount: 0.4, Address: senderWallet.GetAddress()},
	{TxID: "tx4", OutputIndex: 3, Amount: 22.56, Address: recipientWallet.GetAddress()},
	{TxID: "tx5", OutputIndex: 4, Amount: 101.38, Address: recipientWallet.GetAddress()},
	{TxID: "tx6", OutputIndex: 5, Amount: 7.009, Address: senderWallet.GetAddress()},
}

var tr = transaction.CreateTransaction(
	senderWallet.GetPubKey(),
	[]tx.TxOutput{
		{Amount: amount,
			RecipientAddress: recipientWallet.GetAddress(),
		},
	}, utxoSet)

var validOutput = tx.TxOutput{
	RecipientAddress: recipientWallet.GetAddress(),
	Amount:           1.44,
}
