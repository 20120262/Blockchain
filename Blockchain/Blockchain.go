package Blockchain

import (
	"encoding/hex"
	"errors"
)

const genesisCoinbaseData = "The Times 03/Jan/2024 Chancellor on brink of second bailout for banks"

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(transactions, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

}
func CreateBlockchain() *Blockchain {
	tx := Transaction{[]byte(genesisCoinbaseData)}
	return &Blockchain{[]*Block{NewGenesisBlock(&tx)}}
}
func NewBlockChain() *Blockchain {
	return CreateBlockchain()
}
func (bc *Blockchain) GetGenesisBlock() *Block {
	if len(bc.Blocks) > 0 {
		return bc.Blocks[0]
	}
	return nil
}
func (bc *Blockchain) CurrentBlock() *Block {
	if len(bc.Blocks) > 0 {
		return bc.Blocks[len(bc.Blocks)-1]
	}
	return nil
}
func (bc *Blockchain) GetBlock(hash []byte) (*Block, error) {
	for _, block := range bc.Blocks {
		if hex.EncodeToString([]byte(block.Hash)) == hex.EncodeToString(hash) {
			return block, nil
		}
	}
	return nil, errors.New("Block not found")
}
