/*
This is the first version of reading & printing arguments in Go
To run: `$ go run 1.HelloWorld.go Arg1 Arg2`
*/

package main

import (
	"fmt"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
