package tx

import (
	"strconv"

	"github.com/pauldin91/goledger/src/utils"
)

type TxInput struct {
	TxID        string `json:"tx_id"`
	OutputIndex int64  `json:"output_index"`
	Signature   string `json:"signature"`
	PublicKey   string `json:"pubkey"`
}

func (input TxInput) Hash() string {
	var goesIn string = input.TxID + strconv.Itoa(int(input.OutputIndex))
	var hash string = utils.Hash(goesIn)
	return hash
}
