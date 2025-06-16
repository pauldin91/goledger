package pool

import (
	"slices"
	"sync"
	"time"

	"github.com/pauldin91/goledger/src/block"
	"github.com/pauldin91/goledger/src/models"
	"github.com/pauldin91/goledger/src/utils"
)

type MemPool struct {
	Transactions map[string]*models.TransactionDto
	timestamps   map[string]time.Time
	mutex        sync.Mutex
	max_size     int64
}

func NewPool() MemPool {
	return MemPool{
		mutex:        sync.Mutex{},
		Transactions: make(map[string]*models.TransactionDto),
		timestamps:   make(map[string]time.Time),
		max_size:     utils.MemPoolSize,
	}
}

func (p *MemPool) Size() int {
	return len(p.Transactions)
}

func (p *MemPool) AddOrUpdateByID(id string, dto *models.TransactionDto) {
	valid := p.Validate(*dto)
	if dto.Amount > 0 && valid {
		p.mutex.Lock()
		p.Transactions[id] = dto
		p.timestamps[id] = time.Now().UTC()
		p.mutex.Unlock()
	}

}

func (p *MemPool) GetByID(id string) *models.TransactionDto {
	tr, ok := p.Transactions[id]
	if ok {
		return tr
	}
	return nil
}
func (p *MemPool) PurgeExpired() {
	for id, v := range p.Transactions {
		if p.isExpired(*v) {
			p.mutex.Lock()
			delete(p.Transactions, id)
			delete(p.timestamps, id)
			p.mutex.Unlock()
		}
	}
}

func (p *MemPool) isExpired(dto models.TransactionDto) bool {
	stamp, ok := p.timestamps[dto.TxID]
	if ok && stamp.AddDate(0, 0, 2).Unix() < time.Now().UTC().Unix() {
		return true
	}
	return false
}

func (p *MemPool) Validate(tr models.TransactionDto) bool {
	return tr.Amount > 0 && tr.IsValid()
}
func (p *MemPool) Clear() {
	p.Transactions = make(map[string]*models.TransactionDto)
	p.timestamps = make(map[string]time.Time)
}

func (p *MemPool) createBlock(bc *block.Blockchain) {

	var orderedTs = make([]models.TransactionDto, len(p.Transactions))

	p.mutex.Lock()
	for _, c := range p.Transactions {
		orderedTs = append(orderedTs, *c)
	}
	for _, c := range p.Transactions {
		delete(p.Transactions, c.TxID)
		delete(p.timestamps, c.TxID)
	}
	p.mutex.Unlock()

	slices.SortFunc(orderedTs, func(t1, t2 models.TransactionDto) int {
		var result = t1.Timestamp.Compare(t1.Timestamp)
		return result
	})

	var data = models.String(orderedTs)

	bc.AddBlock(data)
}
