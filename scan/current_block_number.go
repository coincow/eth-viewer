package scan

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func GetCurrentBlockNumber(client *ethclient.Client) (*big.Int, *big.Int) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.String())
	return header.Number, header.BaseFee
}
