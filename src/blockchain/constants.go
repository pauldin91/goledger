package blockchain

import (
	"strings"

	"github.com/pauldin91/goledger/src/utils"
)

const (
	MineRate             = 3000
	MiningReward float64 = 3000
)

var GenesisLastHash string = strings.Repeat("0", 64)

var maxByTimestamp = func(k Transaction, t Transaction) Transaction {
	if k.Input.Timestamp.UnixMilli() > t.Input.Timestamp.UnixMilli() {
		return k
	} else {
		return t
	}
}

var findTransactionByAddress = func(t Transaction, a string) bool {
	return t.Input.Address == a
}

var findByAddressAndTimestamp = func(t Transaction, v utils.TimestampAddressFilter) bool {
	_, ex := t.Tx[v.Address]

	return t.Input.Timestamp.After(v.Timestamp) && ex
}

func filter(transaction Transaction) *Transaction {
	var totalOutput float64 = 0.0
	for _, z := range transaction.Tx {
		totalOutput += z.Amount
	}
	if transaction.Input.Amount != totalOutput {
		return nil
	}
	if !transaction.Verify() {
		return nil
	}

	return &transaction

}
