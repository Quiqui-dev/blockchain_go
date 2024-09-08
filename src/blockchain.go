package main

import (
	"bytes"
	"log"
	"time"
)

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock)

	isValidBlock := isBlockValid(newBlock, prevBlock)

	if !isValidBlock {
		log.Fatal("Invalid block")
	}

	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {

	genesisBlock := &Block{
		Index:         0,
		Timestamp:     time.Now().Unix(),
		Data:          []byte("genesis"),
		PrevBlockHash: []byte{},
		Hash:          []byte{},
	}

	genesisBlock.SetHash()

	return genesisBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		[]*Block{
			NewGenesisBlock(),
		},
	}
}

func isBlockValid(newBlock *Block, oldBlock *Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if !bytes.Equal(oldBlock.Hash, newBlock.PrevBlockHash) {
		return false
	}

	return true
}
