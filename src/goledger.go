package goledger

import (
	"github.com/pauldin91/goledger/src/block"
	"github.com/pauldin91/goledger/src/models"
	"github.com/pauldin91/goledger/src/pool"
	"github.com/pauldin91/goledger/src/transaction"
	"github.com/pauldin91/goledger/src/tx"
)

type Block = block.Block
type Blockchain = block.Blockchain
type Transaction = transaction.Transaction
type TransactionDto = models.TransactionDto
type Wallet = transaction.Wallet
type MemPool = pool.MemPool
type TxInput = tx.TxInput
type TxOutput = tx.TxOutput
