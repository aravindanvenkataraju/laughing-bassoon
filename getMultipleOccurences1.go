package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	gets lines as inputs
	stops when input is *
	prints lines that repeat and the number of occurences
*/

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "*" {
			break
		}
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
