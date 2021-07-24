package scan

import (
	"context"
	"eth-viewer2/db"
	"github.com/ethereum/go-ethereum/common"
)

func AddressIsContract(address string) bool {
	isContract, exist := db.EthAddressIsContract(address)
	if exist {
		return isContract
	}

	client := GetClient()
	if nil == client {
		return false
	}

	data, err := client.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if nil != err {
		return false
	}
	if 0 != len(data) {
		db.InsertEthAddress(address, 1, "")
		return true
	} else {
		db.InsertEthAddress(address, 0, "")
		return false
	}
}
