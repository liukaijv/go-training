package main

import (
	"testing"
	"fmt"
	"encoding/hex"
)

func TestNewBlock(t *testing.T) {

	block := NewBlock("new block", []byte{})

	fmt.Println("data:", string(block.Data))
	fmt.Println("hash", hex.EncodeToString(block.Hash))

}
