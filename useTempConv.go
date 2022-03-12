package main

import (
	"fmt"
	"tempconv"
)

func main() {
	fmt.Println(tempconv.AbsoluteZeroC)
	fmt.Printf("%v = %v\n", tempconv.Celsius(20), tempconv.CToF(tempconv.Celsius(20)))
	fmt.Printf("%v = %v\n", tempconv.Kelvin(0), tempconv.KToC(tempconv.Kelvin(20)))
}
