package tests

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/pauldin91/goledger/src/block"
)

func TestCreate(t *testing.T) {
	e := block.Create()
	jsonGen, _ := json.Marshal(genesisBlock)
	jsonFirst, _ := json.Marshal(e.Chain[0])
	if string(jsonFirst) != string(jsonGen) {
		t.Error("First block in chain must be genesis")
	}
}

func TestMineBlock(t *testing.T) {
	for i := 0; i < 64; i++ {
		_ = bc.AddBlock(strconv.Itoa(i))
	}
	time.Sleep(time.Second * 1)
	mined := bc.MineBlock("")
	if !strings.HasPrefix(mined.HashBlock(), strings.Repeat("0", int(1))) {
		t.Errorf("Difficulty  was %s", mined.HashBlock())
	}
}

func TestAddBlock(t *testing.T) {
	e := block.Create()
	jsonMsg, _ := json.Marshal(msg)
	e.AddBlock(string(jsonMsg))

	if len(e.Chain) != 2 {
		t.Error("invalid chain length")
	}
}

func TestReplaceChain(t *testing.T) {
	e := block.Create()
	jsonMsg, _ := json.Marshal(msg)
	e.AddBlock(string(jsonMsg))

	b := block.Create()
	res := e.ReplaceChain(b.Chain)
	if res {
		t.Error("longest chain must not be replaced by smaller ones")
	}
	res = b.ReplaceChain(e.Chain)
	if !res {
		t.Error("smaller chain must be replaced by longer one")
	}

}
