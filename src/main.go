package main

import "fmt"

func main() {

	fmt.Println("Hello World")

	bc := NewBlockchain()

	bc.AddBlock("Test1")
	bc.AddBlock("test2")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %x\n", block.Timestamp)
	}
}
