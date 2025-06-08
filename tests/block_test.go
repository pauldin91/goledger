package tests

import (
	"testing"
)

func TestBlock(t *testing.T) {
	b1 := genesisBlock
	b2 := bc.AddBlock("")
	if b1.GetHash() != b2.GetPrevious() || b2.GetHash() != b2.HashBlock() {
		t.Errorf("1st block's hash %s should match 2nd block's previous hash %s and hash func of 2nd %s should export blocks hash %s\n", b1.GetHash(), b2.GetPrevious(), b2.HashBlock(), b2.GetHash())
	}
}
