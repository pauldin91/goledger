package transaction

func filter(tr Transaction) *Transaction {
	var totalOutput float64 = 0.0
	for _, z := range tr.TxOutputs {
		totalOutput += z.Amount
	}

	if !tr.Verify() {
		return nil
	}

	return &tr

}
