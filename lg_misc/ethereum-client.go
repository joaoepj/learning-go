package lg_misc

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func EthDial() {
	client, err := ethclient.Dial("https://159.89.20.6")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
