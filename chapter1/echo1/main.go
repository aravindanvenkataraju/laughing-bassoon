package main

import (
	"fmt"
	"os"
	"strings"
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

	start = time.Now()
	s = ""
	for _, param := range os.Args[0:] {
		s += sep + param
		sep = " "
	}
	fmt.Println(s)
	elapsed = time.Now().Sub(start)
	fmt.Println("ranged for loop elapsed time: ", elapsed)

	start = time.Now()
	s = strings.Join(os.Args, " ")
	fmt.Println(s)
	elapsed = time.Now().Sub(start)
	fmt.Println("NO for loop elapsed time: ", elapsed)
}
