package tests

import (
	"testing"

	"github.com/pauldin91/goledger/src/utils"
)

func TestCreateTr(t *testing.T) {
	if tr == nil {
		t.Error("unable to create transaction\n")
	}
}

func TestHash(t *testing.T) {
	hh := tr.Hash()

	total := tr.Timestamp.String()
	inputs := ""
	for _, v := range tr.TxInputs {
		inputs += v.Hash()
	}
	outputs := ""
	for _, v := range tr.TxOutputs {
		outputs += v.Hash()
	}
	total += inputs + outputs
	total = utils.Hash(total)
	if hh != total || total != tr.TxID {
		t.Errorf("hash of tx is %s while exported hash is %s \n", tr.TxID, total)
	}

}

func TestSign(t *testing.T) {
	tr.Sign(keyPair)
	isValid := tr.IsValid()
	if !isValid {
		t.Errorf("signatures do not match, output %t\n", tr.IsValid())
	}

}

func TestIsValid(t *testing.T) {
	exported := tr.Hash()
	isValid := utils.VerifySignature(senderWallet.GetPubKey(), []byte(exported), []byte(tr.Signature))

	if isValid != tr.IsValid() {
		t.Errorf("isValid method output %t while export is %t\n", tr.IsValid(), isValid)
	}
}
