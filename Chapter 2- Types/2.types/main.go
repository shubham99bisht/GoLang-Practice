package main

import (
	"fmt"
	"output/tempconv"
)

func main() {
	fmt.Printf("%v\n", tempconv.AbsoluteZeroC.String())
	fmt.Printf("%v\n", tempconv.FreezingC.String())
	fmt.Printf("%v\n", tempconv.BoilingC.String())

	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%v", boilingF.String())
}
