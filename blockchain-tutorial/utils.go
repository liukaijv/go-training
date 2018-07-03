package main

import (
	"bytes"
	"encoding/binary"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

func ReverseBytes(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+i, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}
