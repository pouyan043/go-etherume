package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var infuraURL = "https://mainnet.infura.io/v3/e4a9145f3f5d4bcea7b71028a8cb647d"

// var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL)
	if err != nil {
		log.Fatalf("error to creat a ether client:%v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatalf("error to get a block:%v", err)
	}
	fmt.Println("the block number:", block.Number())

	fmt.Println(block.Number())
	addr := "0x6B21F377C3b0924F12f3A23EDB43e6dD7a61d4BF"
	address := common.HexToAddress(addr)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("error to get the balance:%v", err)
	}
	// fmt.Println("the balance amount:", balance)
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	// fmt.Println("calculated amount:", fbalance)
	BalanceEther := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("the ether amount", BalanceEther)
}
