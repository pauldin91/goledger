package blockchain

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/pauldin91/goledger/src/utils"
)

type Transaction struct {
	Id     uuid.UUID             `json:"id"`
	Input  utils.Data            `json:"input"`
	Tx     map[string]utils.Data `json:"tx"`
	Amount float64               `json:"amount"`
}

func (t Transaction) String() string {
	jsonT, _ := json.Marshal(t)
	return string(jsonT)
}
func transactionWithOutputs(senderWallet Wallet, outputs []utils.Data, amount float64) Transaction {
	transaction := Transaction{
		Id: uuid.New(),
	}
	transaction.Tx = make(map[string]utils.Data)
	for _, o := range outputs {

		transaction.Tx[o.Address] = o
	}
	transaction.Amount = amount
	transaction.Input = utils.CreateData(senderWallet.keyPair.GetPublicKey(), senderWallet.Balance)
	transaction.sign(senderWallet)
	return transaction
}

func (t *Transaction) sign(wallet Wallet) {
	outs, _ := json.Marshal(&t.Tx)
	t.Input.Signature = wallet.keyPair.Sign(utils.Hash(string(outs)))
}

func NewTransaction(senderWallet Wallet, recipient string, amount float64) *Transaction {
	if amount > senderWallet.Balance || amount <= 0 {
		return nil
	}
	outputs := []utils.Data{
		{Amount: senderWallet.Balance - amount, Address: senderWallet.keyPair.GetPublicKey(), Timestamp: time.Now().UTC()},
		{Amount: amount, Address: recipient, Timestamp: time.Now().UTC()},
	}
	var created Transaction = transactionWithOutputs(senderWallet, outputs, amount)
	return &created
}

func (t *Transaction) Update(senderWallet Wallet, recipientAddress string, amount float64) {
	senderOutput := t.Tx[senderWallet.Address]
	if amount > senderOutput.Amount {
		log.Printf("amount %0.8f exceeds balance %0.8f", amount, senderWallet.Balance)
		return
	}
	senderOutput.Amount = senderOutput.Amount - amount
	newlyAdded := utils.CreateData(recipientAddress, amount)
	t.sign(senderWallet)
	t.Tx[newlyAdded.Address] = newlyAdded
}

func (transaction *Transaction) Verify() bool {
	outs, _ := json.Marshal(transaction.Tx)
	var tsString string = utils.Hash(string(outs))
	return utils.VerifySignature(transaction.Input.Address, []byte(tsString), []byte(transaction.Input.Signature))
}

func Reward(minerWallet *Wallet, blockchainWallet *Wallet) *Transaction {
	outputs := []utils.Data{
		{Amount: MiningReward, Address: minerWallet.Address, Timestamp: time.Now().UTC()},
	}
	tr := transactionWithOutputs(*blockchainWallet, outputs, MiningReward)

	return &tr
}
