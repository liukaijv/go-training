package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	client, err := ethclient.Dial("http://127.0.0.1:8545")

	if err != nil {
		fmt.Println(err)
		return
	}

	id, err := client.NetworkID(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("id", id.String())

}
