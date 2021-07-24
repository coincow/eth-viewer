package db

func InsertEthAddress(address string, isContract int, contractName string) {
	db := GetDB()
	if nil == db {
		return
	}

	stmt, err := db.Prepare("INSERT INTO eth_address(address, isContract, cname) values(?,?,?)")
	if err != nil {
		print(err)
		return
	}

	_, err = stmt.Exec(address, isContract, contractName)
	if err != nil {
		print(err)
		return
	}
}

func EthAddressExist(address string) bool {
	return false
}

func EthAddressIsContract(address string) (bool, bool) {
	return false, false
}
