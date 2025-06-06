package utils

import (
	"encoding/json"
	"time"
)

type Data struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    float64   `json:"amount"`
	Address   string    `json:"address"`
	Signature string    `json:"signature"`
}

func CreateData(recipientAddress string, amount float64) Data {
	return Data{
		Timestamp: time.Now().UTC(),
		Amount:    amount,
		Address:   recipientAddress,
	}
}

func (i Data) GetAmount() float64 {
	return i.Amount
}

func (i Data) GetAddress() string {
	return i.Address
}

func (t Data) String() string {
	res, _ := json.Marshal(t)
	return string(res)
}

type TimestampAddressFilter struct {
	Timestamp time.Time
	Address   string
}
