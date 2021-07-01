package scan

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

var mEthClient *ethclient.Client = nil

func GetClient() *ethclient.Client {
	if nil != mEthClient {
		return mEthClient
	}

	client, err := ethclient.Dial("http://47.98.54.147:50002")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	mEthClient = client
	return client
}
