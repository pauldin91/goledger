package pool

import (
	"sync"
	"time"

	"github.com/pauldin91/goledger/src/models"
)

type MemPool struct {
	Transactions map[string]models.TransactionDto
	timestamps   map[string]time.Time
	mutex        sync.Mutex
}

func NewPool() MemPool {
	return MemPool{
		mutex:        sync.Mutex{},
		Transactions: make(map[string]models.TransactionDto),
		timestamps:   make(map[string]time.Time),
	}
}

func (p *MemPool) Size() int {
	return len(p.Transactions)
}

func (p *MemPool) AddOrUpdateByID(id string, dto models.TransactionDto) {
	if dto.Amount > 0 && p.Validate(dto) {
		p.mutex.Lock()
		p.Transactions[id] = dto
		p.timestamps[id] = time.Now().UTC()
		p.mutex.Unlock()
	}

}

func (p *MemPool) GetByID(id string) *models.TransactionDto {
	tr, ok := p.Transactions[id]
	if ok {
		return &tr
	}
	return nil
}
func (p *MemPool) PurgeExpired() {
	for id, v := range p.Transactions {
		if p.isExpired(v) {
			p.mutex.Lock()
			delete(p.Transactions, id)
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
	return tr.Amount > 0 && p.isExpired(tr) && tr.IsValid()
}
func (p *MemPool) Clear() {
	p.Transactions = make(map[string]models.TransactionDto)
	p.timestamps = make(map[string]time.Time)
}
