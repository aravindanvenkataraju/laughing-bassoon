package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	check arguments
	if arguments available, assumed as file name and count occurences in the file
	else, look for entries in std io
*/

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) != 0 {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v \n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	} else {
		countLines(os.Stdin, counts)
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "*" {
			break
		}
		counts[input.Text()]++
	}
}
