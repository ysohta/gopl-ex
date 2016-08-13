package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	err := scanner.Err()
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
	}

	fmt.Print("\nword\tcount\n")
	for s, n := range counts {
		fmt.Printf("%s\t%d\n", s, n)
	}
}
