package scan

import (
	_ "fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func ProcessTransaction(tx *types.Transaction, baseFee *big.Int) {
	chainId := tx.ChainId()
	from := ""
	msg, err := tx.AsMessage(types.NewEIP155Signer(chainId), baseFee)
	if err == nil {
		from = msg.From().Hex()
	} else {
		log.Println(err.Error())
	}
	to := tx.To().String()
	amount := tx.Value().String()
	log.Println("from:" + from + "    to:" + to + "    amount:" + amount)
}
