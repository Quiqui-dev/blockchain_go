package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Index         int64
	Timestamp     int64
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	index := []byte(strconv.FormatInt(b.Index, 10))
	blockHeaders := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp, index}, []byte{})

	hash := sha256.Sum256(blockHeaders)

	b.Hash = hash[:]
}

func NewBlock(data string, PrevBlock *Block) *Block {
	block := &Block{
		Index:         PrevBlock.Index + 1,
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: PrevBlock.Hash,
		Hash:          []byte{},
	}

	block.SetHash()
	return block
}
