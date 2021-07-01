package scan

import "log"

func TestScanCurrent() {
	client := GetClient()
	currentNumber, baseFee := GetCurrentBlockNumber(client)
	log.Println("currentNumber:" + currentNumber.String())

	ProcessBlock(client, currentNumber, baseFee)
}
