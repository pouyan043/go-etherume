package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/ethereum/go-ethereum/accounts/keystore"
)

var (
	url = "https://sepolia.infura.io/v3/e4a9145f3f5d4bcea7b71028a8cb647d"
)

func main() {
	// 	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	// 	_, err := ks.NewAccount("password")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	_, err = ks.NewAccount("password")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// "7153bab3-1788-4920-8202-1a9af988c361"
	// "96c346b4bbb7316dc4fa5ba742365b7d5fcdb906"
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("7153bab3-1788-4920-8202-1a9af988c361")
	a2 := common.HexToAddress("96c346b4bbb7316dc4fa5ba742365b7d5fcdb906")
	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}
	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Balance 1:", b1)
	fmt.Println("Balance 2:", b2)
}
