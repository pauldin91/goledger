package blockchain

import (
	"encoding/json"
	"time"

	"github.com/pauldin91/goledger/src/utils"
)

type Block struct {
	Index      int64     `json:"index"`
	Nonce      int64     `json:"nonce"`
	Difficulty int64     `json:"difficulty"`
	LastHash   string    `json:"previous_block_hash"`
	Hash       string    `json:"current_block_hash"`
	Data       string    `json:"data"`
	Timestamp  time.Time `json:"timestamp"`
}

func Genesis() Block {
	block := Block{
		LastHash: GenesisLastHash,
		Nonce:    0,
	}
	block.Data = ""
	block.Hash = block.GetHash()
	return block
}

func (b *Block) GetHash() string {
	var record string = string(b.Index) + string(b.Nonce) + string(b.Difficulty) + b.LastHash + b.Data + b.Timestamp.Format(time.RFC3339)
	return utils.Hash(record)
}

func (b *Block) ToString() string {
	json, _ := json.Marshal(b)
	return string(json)
}
