package main

import (
	"testing"
	"fmt"
	"encoding/hex"
)

func TestNewBlockChain(t *testing.T) {
	bc := NewBlockChain()

	bc.addBlock("block 1")
	bc.addBlock("block 2")

	for _, block := range bc.blocks {
		fmt.Println("hash", hex.EncodeToString(block.Hash))
		fmt.Println("prevHash", hex.EncodeToString(block.PrevBlockHash))
		fmt.Println("data:", string(block.Data))
	}
}
