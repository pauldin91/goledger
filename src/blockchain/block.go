package blockchain

import (
	"fmt"
	"strings"
	"time"

	"github.com/pauldin91/goledger/src/common"
)

var MineRate = 3000

var GenesisLastHash = strings.Repeat("0", 32)

type Block struct {
	Timestamp  time.Time `json:"timestamp"`
	LastHash   string    `json:"last_hash"`
	Hash       string    `json:"hash"`
	Data       string    `json:"data"`
	Nonce      int64     `json:"nonce"`
	Difficulty int64     `json:"difficulty"`
}

func Genesis() Block {
	block := Block{
		LastHash: GenesisLastHash,
		Nonce:    0,
	}
	block.Data = ""
	block.Hash = common.Hash(block.ToString())
	return block
}

func (b *Block) ToString() string {
	return fmt.Sprintf("Timestamp: %s\nLastHash: %s\nHash: %s\nData: %s\nNonce: %d\nDifficulty: %d\n",
		b.Timestamp.Format(time.RFC3339), b.LastHash, b.Hash, b.Data, b.Nonce, b.Difficulty)
}

func AdjustDifficulty(lastBlock Block, currentTime time.Time) int64 {
	diff := lastBlock.Difficulty
	var start time.Time
	if lastBlock.Timestamp.IsZero() {
		start = time.Now().UTC()
	} else {
		start = lastBlock.Timestamp
	}
	dur := start.UnixMilli() + int64(MineRate)

	if dur > currentTime.UnixMilli() {
		diff += 1
	} else {
		diff -= 1
		if diff <= 0 {
			diff = 1
		}
	}
	return diff
}

func MineBlock(lastBlock Block, data string) Block {

	var hash string
	var timestamp time.Time
	var nonce int64 = 0

	for {
		nonce++
		timestamp = time.Now().UTC()
		difficulty := AdjustDifficulty(lastBlock, timestamp)
		pref := strings.Repeat("0", int(difficulty))
		copy := Block{
			Nonce:      nonce,
			Timestamp:  timestamp,
			Difficulty: difficulty,
			LastHash:   lastBlock.Hash,
			Data:       data,
		}
		hash = common.Hash(copy.ToString())
		copy.Hash = hash
		if strings.HasPrefix(copy.Hash, pref) {
			return copy
		}
	}
}
