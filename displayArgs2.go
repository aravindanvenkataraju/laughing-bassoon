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
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	elapsed := time.Now().Sub(start)
	fmt.Println("traditional for loop elapsed time: ", elapsed)

}
