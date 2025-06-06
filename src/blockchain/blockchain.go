package blockchain

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/pauldin91/goledger/src/utils"
)

type Blockchain struct {
	Chain []Block
}

func (bc *Blockchain) String() string {
	js, _ := json.Marshal(bc.Chain)
	return string(js)

}

func Create() *Blockchain {
	bc := Blockchain{}
	bc.Chain = append(bc.Chain, Genesis())
	return &bc
}

func (bc *Blockchain) AddBlock(data string) Block {
	block := bc.MineBlock(data)
	bc.Chain = append(bc.Chain, block)
	return block
}
func IsValid(bc []Block) bool {

	jsonGenesis, _ := json.Marshal(bc[0])
	gen, _ := json.Marshal(Genesis())
	if string(jsonGenesis) != string(gen) {
		return false
	}
	for i := 1; i < len(bc); i++ {
		block := bc[i]
		lastBlock := bc[i-1]
		expectedHash := block.HashBlock()
		if block.previous != lastBlock.hash ||
			block.hash != expectedHash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) ReplaceChain(newChain []Block) bool {

	isValid := IsValid(newChain)

	if len(newChain) <= len(bc.Chain) || !isValid {
		return false
	}
	bc.Chain = newChain
	return true
}

func (bc *Blockchain) MineBlock(data string) Block {

	var nonce int64 = 0
	var difficulty int64 = 4
	var lastBlock Block
	if len(bc.Chain) == 0 {
		bc.Chain = append(bc.Chain, Genesis())
	}
	lastBlock = bc.Chain[len(bc.Chain)-1]
	if lastBlock.index%2048 == 0 {
		difficulty = utils.AdjustDifficulty(lastBlock.difficulty, lastBlock.timestamp, time.Now().UTC(), MineRate)
	}
	for {
		nonce++
		pref := strings.Repeat("0", int(difficulty))
		copy := lastBlock.Create(nonce, difficulty, data)
		if strings.HasPrefix(copy.hash, pref) {
			return copy
		}
	}
}
