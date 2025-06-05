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
	block := bc.MineBlock(data, true)
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
		expectedHash := block.GetHash()
		if block.LastHash != lastBlock.Hash ||
			block.Hash != expectedHash {
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

func (bc *Blockchain) MineBlock(data string, adjusting bool) Block {

	var timestamp time.Time
	var nonce int64 = 0
	var difficulty int64 = 4
	var lastBlock Block
	if len(bc.Chain) == 0 {
		bc.Chain = append(bc.Chain, Genesis())
	}
	lastBlock = bc.Chain[len(bc.Chain)-1]
	for {
		nonce++
		timestamp = time.Now().UTC()
		if adjusting {
			difficulty = utils.AdjustDifficulty(lastBlock.Difficulty, lastBlock.Timestamp, timestamp, MineRate)
		}

		pref := strings.Repeat("0", int(difficulty))
		copy := Block{
			Nonce:      nonce,
			Timestamp:  timestamp,
			Difficulty: difficulty,
			LastHash:   lastBlock.Hash,
			Data:       data,
		}
		copy.Hash = copy.GetHash()
		if strings.HasPrefix(copy.Hash, pref) {
			return copy
		}
	}
}
