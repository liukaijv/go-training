package rsa

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func loadKeys() (publicKey, privateKey []byte, err error) {

	publicKey, err = ioutil.ReadFile("public.pem")
	if err != nil {
		return
	}

	privateKey, err = ioutil.ReadFile("private.pem")
	if err != nil {
		return
	}

	return

}

//test rsa加解密
func TestRsa(t *testing.T) {

	publicKey, privateKey, err := loadKeys()

	if err != nil {
		t.Fatal(err)
	}

	originText := "test rsa"
	cipherData, err := Encrypt([]byte(originText), publicKey)
	if err != nil {
		t.Fatal(err)
	}
	originData, err := Decrypt(cipherData, privateKey)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(originData))
	if originText != string(originData) {
		t.Fatal("rsa加解密失败")
	}

}

func TestGenKey(t *testing.T) {
	bits := 1024
	err := GenKey(bits)
	if err != nil {
		t.Fatal(err)
	}
}
