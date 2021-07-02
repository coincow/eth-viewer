package main

import (
	"eth-viewer/scan"
	"eth-viewer/test"
	"log"
	"time"
)

func main() {
	address := "0x4363A7f50D7d91a791D1dc5971c83632D29196f0"
	balance := test.GetBalance(address)
	log.Println("address:" + address + "-------balance:" + balance.String())

	scan.TestScanCurrent()
	time.Sleep(2 * time.Second)
}
