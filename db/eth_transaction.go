package db

import "time"

func InsertEthTrans(txHash string, blockNum int64, timestamp time.Time, from string, to string, value int64, fee int64, gasPrice int64) {
	db := GetDB()
	if nil == db {
		return
	}

	stmt, err := db.Prepare("INSERT INTO eth_transaction(txhash, blocknum, timestamp1, from1, to1, value1, fee, gasprice) values(?,?,?,?,?,?,?,?)")
	if err != nil {
		print(err)
		return
	}

	_, err = stmt.Exec(txHash, blockNum, timestamp.Unix(), from, to, value, fee, gasPrice)
	if err != nil {
		print(err)
		return
	}
}
