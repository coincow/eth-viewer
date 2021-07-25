package scan

import (
	"eth-viewer/db"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"time"
)

func ProcessTransaction(tx *types.Transaction, baseFee *big.Int, time time.Time, blockNumber int64) {
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
	db.InsertEthTrans(tx.Hash().String(), blockNumber, time, from, to, tx.Value().Int64(), 1, tx.GasPrice().Int64())
}
