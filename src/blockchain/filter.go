package blockchain

import "github.com/pauldin91/goledger/src/utils"

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
	_, ex := t.Output[v.Address]

	return t.Input.Timestamp.After(v.Timestamp) && ex
}
