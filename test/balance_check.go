package test

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetBalance(address string) big.Int {

	client, err := ethclient.Dial("http://47.98.54.147:50002")
	if err != nil {
		log.Println(err)
		return big.Int{}
	}

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Println(err)
		return big.Int{}
	}
	fmt.Println(balance)
	return *balance
}
