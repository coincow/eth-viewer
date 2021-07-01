package scan

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func ProcessBlock(client *ethclient.Client, blockNumber *big.Int, baseFee *big.Int) {
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		ProcessTransaction(tx, baseFee)
	}
}
