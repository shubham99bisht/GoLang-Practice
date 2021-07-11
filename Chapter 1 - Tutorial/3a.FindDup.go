/*
FindDup finds duplicate lines from in the input from stdin
To run:
 `$ go run 3.FindDup.go`
 Enter multiple lines in input
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
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}
