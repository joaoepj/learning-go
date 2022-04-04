package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/fdlimit"
)


func main() {
	limit, err := fdlimit.Maximum()
	fmt.Println(limit, err)

}
