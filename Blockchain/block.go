package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Transactions  []*Transaction
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) setHash() {
	var headers []byte
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers = append(headers, timestamp...)
	for _, transaction := range b.Transactions {
		headers = append(headers, transaction.Data...)
	}
	headers = append(headers, b.PrevBlockHash...)
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}}
	block.setHash()
	return block

}

func NewGenesisBlock(coinbase *Transaction) *Block {

	return NewBlock([]*Transaction{coinbase}, []byte{})

}

func (b *Block) HashTransaction() []byte {
	var txHashes [][]byte
	var txHash [32]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Data)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}
