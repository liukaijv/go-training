package main

import (
	"bytes"
	"context"
	"ethereum-develop/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func main() {

	privateKeyHash := "79e66c43f3eb1ad710321228ddd13229f3830495b1f8068bf21d3f66cc7f8ff8"
	rawUrl := "http://127.0.0.1:8545"

	conn, err := ethclient.Dial(rawUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	pk, err := crypto.HexToECDSA(privateKeyHash)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}
	auth := bind.NewKeyedTransactor(pk)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	// Deploy a new awesome contract for the binding demo
	address, tx, token, err := contracts.DeployTTToken(auth, conn)
	if err != nil {
		log.Fatalf("Failed to deploy new token contract: %v", err)
	}
	fmt.Printf("Contract pending deploy: 0x%x\n", address)
	fmt.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
	startTime := time.Now()
	fmt.Printf("TX start @:%s", time.Now())

	ctx := context.Background()
	addressAfterMined, err := bind.WaitDeployed(ctx, conn, tx)
	if err != nil {
		log.Fatalf("failed to deploy contact when mining :%v", err)
	}
	fmt.Printf("tx mining take time:%s\n", time.Now().Sub(startTime))
	if bytes.Compare(address.Bytes(), addressAfterMined.Bytes()) != 0 {
		log.Fatalf("mined address :%s,before mined address:%s", addressAfterMined, address)
	}

	name, err := token.Name(&bind.CallOpts{Pending: true})
	if err != nil {
		log.Fatalf("Failed to retrieve pending name: %v", err)
	}
	fmt.Println("Pending name:", name)
}
