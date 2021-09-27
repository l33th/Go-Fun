package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	conn, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatalf("Sorry, something went wrong!", err)
	}
	ctx := context.Background()
	tx, pending, _ := conn.TransactionByHash(ctx, common.HexToHash("Some Hash from Etherscan site"))
	if (!pending) {
		fmt.Println(tx)
	}
}
