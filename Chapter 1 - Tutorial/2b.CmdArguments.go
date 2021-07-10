/*
This is second version of Cmd line arguments in Go using range in for loop
To run `$ go run 2a.CmdArguments.go Arg1 Arg2 Arg3`
Expected Output: 'Arg1 Arg2 Arg3'
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
