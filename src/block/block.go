package block

import (
	"encoding/json"
	"strconv"
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
		index:      0,
		previous:   utils.GenesisLastHash,
		Nonce:      0,
		difficulty: 1,
		data:       "",
	}
	block.hash = block.HashBlock()
	return block
}

func (b *Block) Create(nonce int64, diff int64, data string) Block {
	var created Block = Block{
		index:      b.index + 1,
		previous:   b.hash,
		data:       data,
		timestamp:  time.Now().UTC(),
		Nonce:      nonce,
		difficulty: diff,
	}
	created.hash = created.HashBlock()
	return created

}
func (b Block) GetHash() string {
	return b.hash
}
func (b Block) GetPrevious() string {
	return b.previous
}
func (b Block) HashBlock() string {
	var record string = strconv.FormatInt(b.index, 10) + strconv.FormatInt(b.Nonce, 10) +
		strconv.FormatInt(b.difficulty, 10) +
		b.previous + b.data + b.timestamp.Format(time.RFC3339)

	resutl := utils.Hash(record)
	return resutl
}

func (b *Block) ToString() string {
	json, _ := json.Marshal(b)
	return string(json)
}
