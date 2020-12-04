package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewBlock(t *testing.T) {

	block := NewBlock("new block", []byte{})

	fmt.Println("data:", string(block.Data))
	fmt.Println("hash", hex.EncodeToString(block.Hash))

}
