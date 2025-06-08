package transaction

type TxInput struct {
	TxID        string `json:"tx_id"`
	OutputIndex int64  `json:"output_index"`
	Signature   string `json:"signature"`
	PublicKey   string `json:"pubkey"`
}
