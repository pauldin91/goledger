package tests

import (
	"github.com/pauldin91/goledger/src/block"
	"github.com/pauldin91/goledger/src/pool"
	"github.com/pauldin91/goledger/src/transaction"
	"github.com/pauldin91/goledger/src/tx"
)

const (
	recipientAddress string = "r3ciP13nT4Ddr3$5"
)

var genesisBlock = block.Genesis()

var tp = pool.MemPool{}

var amount float64 = 10.0

var bc = block.Create()

var wallet = transaction.NewWallet()

var utxoSet []tx.UTXO

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

var tr = transaction.CreateTransaction(
	wallet.GetPubKey(),
	[]tx.TxOutput{
		{Amount: amount,
			RecipientAddress: recipientAddress,
		},
	}, utxoSet)

var msg = tx.TxOutput{
	RecipientAddress: "r3ciP13nT",
	Amount:           50.44,
}
