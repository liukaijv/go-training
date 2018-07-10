package aes

import (
	"testing"
	"encoding/base64"
	"fmt"
)

func TestAes(t *testing.T) {

	originStr := "test aes"

	key := []byte("aaaabbbbccccdddd")
	result, err := Encrypt([]byte(originStr), key)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(result))
	decryptData, err := Decrypt(result, key)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(decryptData))

	if string(decryptData) != originStr {
		t.Fatal("aes fail")
	}

}
