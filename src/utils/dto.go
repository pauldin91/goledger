package utils

import (
	"encoding/json"
	"time"
)

type Input struct {
	Timestamp time.Time `json:"timestamp"`
	Amount    float64   `json:"amount"`
	Address   string    `json:"address"`
	Signature string    `json:"signature"`
}

func (i Input) GetAmount() float64 {
	return i.Amount
}

func (i Input) GetAddress() string {
	return i.Address
}

func (t Input) String() string {
	res, _ := json.Marshal(t)
	return string(res)
}

type TimestampAddressFilter struct {
	Timestamp time.Time
	Address   string
}
