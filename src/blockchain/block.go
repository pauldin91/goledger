package blockchain

import (
	"encoding/json"
	"time"

	"github.com/pauldin91/goledger/src/utils"
)

type Block struct {
	Timestamp  time.Time `json:"timestamp"`
	LastHash   string    `json:"previous_block_hash"`
	Hash       string    `json:"current_block_hash"`
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
	block.Hash = utils.Hash(block.ToString())
	return block
}

func (b *Block) ToString() string {
	json, _ := json.Marshal(b)
	return string(json)
}
