/*
FindDupiInFile finds duplicate lines in a file which is provided as cmd line argument
To run: `$ go run 3.FindDupInFile.go filePath`
Expected output: '<duplicate line> <count>'
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
