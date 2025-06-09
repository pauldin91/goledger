package tx

type UTXO struct {
	TxID        string  `json:"txid"`
	OutputIndex int64   `json:"output_index"`
	Amount      float64 `json:"amount"`
	Address     string  `json:"address"`
}

func (utxo UTXO) Map() TxInput {
	return TxInput{
		TxID:        utxo.TxID,
		OutputIndex: utxo.OutputIndex,
	}
}
