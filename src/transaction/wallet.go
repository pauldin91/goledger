package transaction

import (
	"github.com/pauldin91/goledger/src/models"
	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

type Wallet struct {
	keyPair      *utils.KeyPair
	address      string
	transactions chan models.TransactionDto
}

func NewWallet() Wallet {
	var wallet Wallet = Wallet{
		keyPair: utils.NewKeyPair(),
	}
	wallet.address = utils.Hash(wallet.keyPair.GetPublicKey())
	return wallet
}

func (w Wallet) CalculateBalance(utxoSet map[string]tx.TxOutput) float64 {
	balance := 0.0
	for _, output := range utxoSet {
		if output.RecipientAddress == w.address {
			balance += output.Amount
		}
	}
	return balance
}

func (w Wallet) Send(newOutput tx.TxOutput, utxoSet map[string]tx.TxOutput) bool {
	if newOutput.Amount <= 0.0 {
		return false
	}
	balance := w.CalculateBalance(utxoSet)
	if balance >= newOutput.Amount {
		tr := NewTransaction(newOutput)
		w.transactions <- tr.Map()
		return true
	} else {
		return false
	}
	return false
}
