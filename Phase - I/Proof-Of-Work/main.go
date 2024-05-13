package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		pow := NewProofOfWork(block)
		fmt.Printf("Prev. hash: %s\n", hex.EncodeToString(block.PrevBlockHash))
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n\n", hex.EncodeToString(block.Hash))
		fmt.Printf("POW: %s", strconv.FormatBool(pow.validate()))
		fmt.Println()
	}
}
