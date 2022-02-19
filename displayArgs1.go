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
	s := strings.Join(os.Args, " ")
	fmt.Println(s)
	elapsed := time.Now().Sub(start)
	fmt.Println("NO for loop elapsed time: ", elapsed)
}
