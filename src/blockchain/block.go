package blockchain

import (
	"encoding/json"
	"time"

	"github.com/pauldin91/goledger/src/utils"
)

type Block struct {
	index      int64
	Nonce      int64
	difficulty int64
	previous   string
	hash       string
	data       string
	timestamp  time.Time
}

func Genesis() Block {
	block := Block{
		index:    0,
		previous: GenesisLastHash,
		Nonce:    0,
	}
	block.data = ""
	block.hash = block.GetHash()
	return block
}

func (b *Block) Create(nonce int64, diff int64, data string) Block {
	return Block{
		index:      b.index + 1,
		previous:   b.hash,
		data:       data,
		timestamp:  time.Now().UTC(),
		Nonce:      nonce,
		difficulty: diff,
	}

}

func (b *Block) GetHash() string {
	var record string = string(b.index) + string(b.Nonce) + string(b.difficulty) + b.previous + b.data + b.timestamp.Format(time.RFC3339)
	return utils.Hash(record)
}

func (b *Block) ToString() string {
	json, _ := json.Marshal(b)
	return string(json)
}
