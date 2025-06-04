package tests

import (
	"strings"
	"testing"
	"time"

	"github.com/pauldin91/goledger/src/blockchain"
	"github.com/pauldin91/goledger/src/utils"
)

func TestGenesis(t *testing.T) {

	if len(genesisBlock.Data) != 0 ||
		genesisBlock.LastHash != blockchain.GenesisLastHash {
		t.Error("Data and lasthash should be empty for genesis")
	}
	if genesisBlock.Difficulty != 0 ||
		genesisBlock.Nonce != 0 {
		t.Error("Difficulty and nonce should be 0 for genesis")
	}
	block := blockchain.Block{
		LastHash: blockchain.GenesisLastHash,
		Nonce:    0,
	}
	block.Data = ""
	block.Hash = utils.Hash(block.ToString())
	if genesisBlock.Hash != block.Hash {
		t.Error("Hashes missmatch")
	}

	if !genesisBlock.Timestamp.IsZero() {
		t.Error("Genesis time is not zero")
	}
}

func TestAdjustDifficulty(t *testing.T) {
	diff := utils.AdjustDifficulty(genesisBlock.Difficulty, genesisBlock.Timestamp, time.Now().UTC(), 1000)
	if diff != 1 {
		t.Errorf("Difficulty should be %d\n", diff)
	}
	genesisBlock.Difficulty = 5
	diff = utils.AdjustDifficulty(genesisBlock.Difficulty, genesisBlock.Timestamp, time.Now().UTC().Add(time.Duration(time.Second*4)), 1000)
	if diff != 4 {
		t.Errorf("Difficulty should be %d\n", diff)
	}
	genesisBlock.Difficulty = 0
}

func TestMineBlock(t *testing.T) {
	mined := bc.MineBlock("", false)
	if !strings.HasPrefix(mined.Hash, strings.Repeat("0", int(genesisBlock.Difficulty))) {
		t.Errorf("Difficulty was %d while output was %s", genesisBlock.Difficulty, mined.Hash)
	}
	genesisBlock.Difficulty = 5
	time.Sleep(time.Second * 4)
	mined = bc.MineBlock("", false)
	if !strings.HasPrefix(mined.Hash, strings.Repeat("0", int(genesisBlock.Difficulty-1))) {
		t.Errorf("Difficulty was %d while output was %s", genesisBlock.Difficulty, mined.Hash)
	}
	genesisBlock.Difficulty = 0
}
