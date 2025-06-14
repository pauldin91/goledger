package transaction

import (
	"github.com/pauldin91/goledger/src/pool"
	"github.com/pauldin91/goledger/src/tx"
	"github.com/pauldin91/goledger/src/utils"
)

type Wallet struct {
	keyPair utils.KeyPair
	address string
	utxoSet []tx.UTXO
}

func NewWallet() Wallet {
	var wallet Wallet = Wallet{
		keyPair: utils.NewKeyPair(),
		utxoSet: make([]tx.UTXO, 0),
	}
	wallet.address = utils.Hash(wallet.keyPair.GetPublicKey())
	return wallet
}

func NewWalletWithKeys(keyPair utils.KeyPair) Wallet {
	var wallet Wallet = Wallet{
		keyPair: keyPair,
		utxoSet: make([]tx.UTXO, 0),
	}
	wallet.address = utils.Hash(wallet.keyPair.GetPublicKey())
	return wallet
}

func (w *Wallet) WithUTXOs(utxos []tx.UTXO) {
	w.utxoSet = utxos
}

func (w Wallet) GetAddress() string {
	return w.address
}
func (w Wallet) GetPubKey() string {
	return w.keyPair.GetPublicKey()
}

func (w Wallet) CalculateBalance() float64 {
	balance := 0.0
	for _, output := range w.utxoSet {
		if output.Address == w.address {
			balance += output.Amount
		}
	}
	return balance
}

func (w Wallet) Send(recipient tx.TxOutput, pool *pool.MemPool) bool {
	if recipient.Amount <= 0.0 {
		return false
	}
	balance := w.CalculateBalance()
	if balance >= recipient.Amount {
		outputs, selectedUTXOs := w.GetTransactionUTXOs(recipient)
		tr := CreateTransaction(w.keyPair.GetPublicKey(), outputs, selectedUTXOs)
		tr.Sign(w.keyPair)
		mapped := tr.Map()
		pool.AddOrUpdateByID(tr.Hash(), &mapped)
		return true
	} else {
		return false
	}
}

func (w Wallet) GetTransactionUTXOs(recipient tx.TxOutput) ([]tx.TxOutput, []tx.UTXO) {
	var selectedUTXOs []tx.UTXO
	var total float64
	outputs := []tx.TxOutput{
		recipient,
	}

	for _, utxo := range w.utxoSet {
		selectedUTXOs = append(selectedUTXOs, utxo)
		total += utxo.Amount
		if total >= recipient.Amount {
			break
		}
	}

	change := total - recipient.Amount
	if change > 0 {
		outputs = append(outputs, tx.TxOutput{
			Amount:           change,
			RecipientAddress: w.address,
		})
	}

	return outputs, selectedUTXOs

}
