package tests

import (
	"time"

	"github.com/pauldin91/goledger/src/blockchain"
	"github.com/pauldin91/goledger/src/utils"
)

var genesisBlock = blockchain.Genesis()
var pool = blockchain.TransactionPool{}
var amount float64 = 10.0
var tp = blockchain.TransactionPool{}
var bc = blockchain.Create()
var senderWallet = blockchain.NewWallet(100.0)
var recipientWallet = blockchain.NewWallet(0.0)

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

var transaction = blockchain.NewTransaction(senderWallet, recipientWallet.Address, amount)

var msg = utils.Data{
	Timestamp: time.Now().UTC(),
	Address:   "r3ciP13nT",
	Amount:    50.44,
}
