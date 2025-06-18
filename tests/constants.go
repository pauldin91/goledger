package tests

import (
	"github.com/pauldin91/goledger/src/block"
	"github.com/pauldin91/goledger/src/pool"
	"github.com/pauldin91/goledger/src/transaction"
	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

const (
	amount float64 = 1.0
)

var transmitTsChan chan string = make(chan string)
var doneChan chan bool = make(chan bool)

var keyPair = utils.NewKeyPair()
var genesisBlock = block.Genesis()

var tpool = pool.NewPool(transmitTsChan, doneChan)

var bc = block.Blockchain{
	Chain: []block.Block{genesisBlock},
}

var senderWallet = transaction.NewWalletWithKeys(keyPair)
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
	{TxID: "tx2", OutputIndex: 1, Amount: 42.0, Address: senderWallet.GetAddress()},
	{TxID: "tx3", OutputIndex: 2, Amount: 5.4, Address: senderWallet.GetAddress()},
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

var trsForPool = []*transaction.Transaction{
	transaction.CreateTransaction(senderWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount + 0.1, RecipientAddress: recipientWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(recipientWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount - 0.2, RecipientAddress: senderWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(senderWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount + 0.3, RecipientAddress: recipientWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(recipientWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount - 0.4, RecipientAddress: senderWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(senderWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount + 0.5, RecipientAddress: recipientWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(recipientWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount - 0.6, RecipientAddress: senderWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(senderWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount + 0.7, RecipientAddress: recipientWallet.GetAddress()}}, utxoSet),
	transaction.CreateTransaction(recipientWallet.GetPubKey(),
		[]tx.TxOutput{{Amount: amount - 0.8, RecipientAddress: senderWallet.GetAddress()}}, utxoSet),
}

var validOutput = tx.TxOutput{
	RecipientAddress: recipientWallet.GetAddress(),
	Amount:           1.44,
}
