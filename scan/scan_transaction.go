package scan

import (
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

	fromContract := AddressIsContract(from)
	toContract := AddressIsContract(to)

	var fromTag = "[0]"
	if fromContract {
		fromTag = "[1]"
	} else {
		fromTag = "[0]"
	}
	var toTag = "[0]"
	if toContract {
		toTag = "[1]"
	} else {
		toTag = "[0]"
	}

	log.Println("from:" + from + fromTag + "    to:" + to + toTag + "    amount:" + amount)
}
