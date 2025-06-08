package transaction

type TxOutput struct {
	Amount  float64 `json:"amount"`
	Address string  `json:"address"`
}
