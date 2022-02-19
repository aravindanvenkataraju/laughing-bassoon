package main

import (
	"fmt"
	"os"
	"time"
)

/*
Program displays all arguments including the executable
*/
func main() {
	start := time.Now()
	var s, sep string
	for _, param := range os.Args[0:] {
		s += sep + param
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Now().Sub(start)
	fmt.Println("ranged for loop elapsed time: ", elapsed)
}
