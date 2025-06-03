package blockchain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pauldin91/goledger/src/utils"
)

type Wallet struct {
	Balance float64 `json:"balance"`
	keyPair *utils.KeyPair
	Address string `json:"address"`
}

func (w Wallet) String() string {
	jsonWallet, _ := json.Marshal(w)
	return string(jsonWallet)
}

func NewWallet(init float64) Wallet {
	res := Wallet{
		Balance: init,
		keyPair: utils.NewKeyPair(),
	}
	res.Address = res.keyPair.GetPublicKey()
	return res
}

func (w Wallet) ToString() string {
	return fmt.Sprintf("Wallet - \npublicKey\t: %s\nbalance: %.8f\n", w.keyPair.GetPublicKey(), w.Balance)
}

func (w Wallet) CalculateBalance(chain Blockchain) float64 {
	var totalTransactions []Transaction
	balance := w.Balance
	for _, b := range chain.Chain {
		var transactions []Transaction
		_ = json.Unmarshal([]byte(b.Data), &transactions)
		totalTransactions = append(totalTransactions, transactions...)
	}
	walletInputTs := utils.FilterBy(totalTransactions, w.keyPair.GetPublicKey(), findTransactionByAddress)

	var start time.Time
	if len(walletInputTs) > 0 {
		recentInputT := utils.Aggregate(walletInputTs, maxByTimestamp)
		balance = recentInputT.Output[w.keyPair.GetPublicKey()].Amount
		start = recentInputT.Input.Timestamp
	}

	v := utils.TimestampAddressFilter{
		Timestamp: start,
		Address:   w.keyPair.GetPublicKey(),
	}

	filteredOutputs := make(map[string]utils.Input)
	utils.SelectMany(totalTransactions, &filteredOutputs, func(t *Transaction, m *map[string]utils.Input) {
		for _, i := range t.Output {
			if i.Address == v.Address && t.Input.Timestamp.After(v.Timestamp) {
				filteredOutputs[i.Address] = i
			}
		}
	})

	for _, b := range filteredOutputs {
		balance += b.Amount
	}

	return balance
}

func (w Wallet) CreateTransaction(recipient string, amount float64, blockchain Blockchain, pool *TransactionPool) bool {

	w.Balance = w.CalculateBalance(blockchain)
	if amount > w.Balance || amount <= 0.0 {
		return false
	} else {
		transaction := NewTransaction(w, recipient, amount)
		pool.AddOrUpdateById(*transaction)
		return true
	}

}
