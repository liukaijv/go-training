package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {

	client, err := rpc.Dial("http://127.0.0.1:8545")

	if err != nil {
		fmt.Println(err)
		return
	}

	var version string
	err = client.Call(&version, "net_version")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("version", version)

}
