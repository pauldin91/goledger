package tests

import (
	"testing"
	"time"
)

func TestCreatePool(t *testing.T) {
	if &tpool == nil {
		t.Error("null mem pool")
	}

}

func TestAddToThePool(t *testing.T) {
	tr.Sign(keyPair)
	mappedDto := tr.Map()

	tpool.AddOrUpdateByID(tr.TxID, &mappedDto)
	if tpool.Size() == 0 {
		t.Error("mempool does not take in add\n")
	}
}

func TestUpdateToThePool(t *testing.T) {
	mappedDto := tr.Map()
	mappedDto.Amount = 1000.0
	tpool.AddOrUpdateByID(tr.TxID, &mappedDto)
	if mappedDto.Amount != 1000.0 {
		t.Errorf("mempool does not update the transaction, output %f.10 while actual should be %f.10 \n", mappedDto.Amount, 1000.0)

	}
}
func TestGetFromPool(t *testing.T) {

	fetched := tpool.GetByID(tr.TxID)
	if fetched.Amount != 1000.0 {
		t.Errorf("mempool does not update the transaction, output %f.10 while actual should be %f.10 \n", fetched.Amount, 1000.0)

	}

}

func TestPurgeExpired(t *testing.T) {
	tr.Sign(keyPair)
	mappedDto := tr.Map()

	tpool.AddOrUpdateByID(tr.TxID, &mappedDto)

	fetched := tpool.GetByID(tr.TxID)
	fetched.Timestamp = time.Now().UTC().AddDate(0, 0, -3)
	tpool.AddOrUpdateByID(fetched.TxID, fetched)
	tpool.PurgeExpired()
	if tpool.Size() != 1 {
		t.Errorf("mempool purges tempered, size is %d", tpool.Size())

	}
	fetched.Timestamp = time.Now().UTC().AddDate(0, 0, 3)
	tpool.AddOrUpdateByID(fetched.TxID, fetched)
}

func TestValidate(t *testing.T) {
	mapped := tr.Map()
	if !tpool.Validate(mapped) {
		t.Error("should validate a valid transaction\n")
	}
}

func TestValidTxs(t *testing.T) {

}
