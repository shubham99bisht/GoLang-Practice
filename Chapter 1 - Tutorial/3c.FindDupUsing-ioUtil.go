/*
This file reads a file entirely into memory and
then processes it line by line to find duplicates
To run: `$ go run 3.FindDupInFile.go filePath`
Expected output: '<count> <duplicate line>'
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	fileName := os.Args[1]
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	for _, line := range strings.Split(string(data), "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
