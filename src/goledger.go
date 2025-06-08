package goledger

import (
	"github.com/pauldin91/goledger/src/block"
	"github.com/pauldin91/goledger/src/transaction"
)

type Block = block.Block
type Blockchain = block.Blockchain
type Transaction = transaction.Transaction
type TransactionPool = transaction.TransactionPool
type TxInput = transaction.TxInput
type TxOutput = transaction.TxOutput
