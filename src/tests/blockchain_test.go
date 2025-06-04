package tests

import (
	"encoding/json"
	"testing"

	"github.com/pauldin91/goledger/src/blockchain"
)

func TestCreate(t *testing.T) {
	e := blockchain.Create()
	jsonGen, _ := json.Marshal(genesisBlock)
	jsonFirst, _ := json.Marshal(e.Chain[0])
	if string(jsonFirst) != string(jsonGen) {
		t.Error("First block in chain must be genesis")
	}
}

func TestAddBlock(t *testing.T) {
	e := blockchain.Create()
	jsonMsg, _ := json.Marshal(msg)
	e.AddBlock(string(jsonMsg))

	if len(e.Chain) != 2 {
		t.Error("invalid chain length")
	}
}

func TestReplaceChain(t *testing.T) {
	e := blockchain.Create()
	jsonMsg, _ := json.Marshal(msg)
	e.AddBlock(string(jsonMsg))

	b := blockchain.Create()
	res := e.ReplaceChain(b.Chain)
	if res {
		t.Error("longest chain must not be replaced by smaller ones")
	}
	res = b.ReplaceChain(e.Chain)
	if !res {
		t.Error("smaller chain must be replaced by longer one")
	}

}
