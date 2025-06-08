package tx

import (
	"strconv"

	"github.com/pauldin91/goledger/src/utils"
)

type TxOutput struct {
	Amount           float64 `json:"amount"`
	RecipientAddress string  `json:"address"`
}

func (input TxOutput) Hash() string {
	var goesIn string = strconv.FormatFloat(input.Amount, 'f', -1, 64) + input.RecipientAddress
	var hash string = utils.Hash(goesIn)
	return hash
}
