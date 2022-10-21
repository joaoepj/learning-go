package lg_misc

import "fmt"

func MaxInt32() int32 {
	return 1<<31 - 1
}

func MaxInt64() int64 {
	return 1<<63 - 1
}

func XorTest() {
	var x, y int32 = 2, 10
	fmt.Printf("x= dec:%v bin:%b, y=  dec:%v bin:%b , xor = dec:%v bin:%b\n", x, x, y, y, (x ^ y), (x ^ y))
}
