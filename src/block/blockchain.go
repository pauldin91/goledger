package block

import (
	"encoding/json"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/pauldin91/goledger/src/utils"
)

type Blockchain struct {
	transmitTsChan chan string
	errorChan      chan error
	doneChan       chan bool
	wg             *sync.WaitGroup
	Chain          []Block
}

func Create(transmitTsChan chan string, doneChan chan bool) *Blockchain {
	bc := Blockchain{
		transmitTsChan: transmitTsChan,
		errorChan:      make(chan error),
		doneChan:       doneChan,
		wg:             &sync.WaitGroup{},
	}
	bc.Chain = append(bc.Chain, Genesis())
	go func() {
		bc.listenToMempool()
	}()
	return &bc
}

func (bc *Blockchain) AddBlock(data string) Block {
	block := bc.mine(data)
	bc.Chain = append(bc.Chain, block)
	return block
}

func (bc *Blockchain) ReplaceChain(newChain []Block) bool {

	isValid := isValid(newChain)

	if len(newChain) <= len(bc.Chain) || !isValid {
		return false
	}
	bc.Chain = newChain
	return true
}

func (bc *Blockchain) mine(data string) Block {

	var nonce int64 = 0
	var difficulty int64 = 4
	var lastBlock Block
	if len(bc.Chain) == 0 {
		bc.Chain = append(bc.Chain, Genesis())
	}
	lastBlock = bc.Chain[len(bc.Chain)-1]
	if lastBlock.index != 0 && lastBlock.index%utils.AdjustDifficultyEvery == 0 {
		difficulty = utils.AdjustDifficulty(lastBlock.difficulty, lastBlock.timestamp, time.Now().UTC(), utils.MineRate)
	}
	pref := strings.Repeat("0", int(difficulty))
	for {
		nonce++
		copy := lastBlock.Create(nonce, difficulty, data)
		if strings.HasPrefix(copy.hash, pref) {
			return copy
		}
	}
}

func isValid(bc []Block) bool {

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

func (bc *Blockchain) listenToMempool() {

	for {
		select {
		case data := <-bc.transmitTsChan:
			bc.AddBlock(data)
			bc.doneChan <- true
		case err := <-bc.errorChan:
			log.Printf("error: %v", err)
			return
		}
	}

}
