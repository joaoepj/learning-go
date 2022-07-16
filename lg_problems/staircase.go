package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// HackerHank Staircase Problem

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for line := 1; int32(line) <= n; line++ {
		stairline(int32(line), n)
	}

}

func stairline(a int32, b int32) {
	var chars []byte

	for i := 0; int32(i) < (b - a); i++ {
		chars = append(chars, ' ')
	}
	for i := 0; int32(i) < a; i++ {
		chars = append(chars, '#')
	}
	fmt.Println(string(chars))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
