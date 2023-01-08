package lg_misc

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/fdlimit"
)

// Just curious about some issue on github Ethereum
// Testing the Limit for File Descriptors
func TestFdLimit() {
	limit, err := fdlimit.Maximum()
	fmt.Println(limit, err)

}
