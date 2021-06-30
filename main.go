package main

import (
	"eth-viewer2/test"
	"log"
)

func main() {
	address := "0xCC5d22bB804D7BFFc48BDAe7563Ba7843FE4928a"
	balance := test.GetBalance(address)
	log.Println("address:" + address + "-------balance:" + balance.String())
}
