// ex04-dup prints duplicated lines in the multiple files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type set map[string]bool

func main() {
	// key : line
	// value : set of files (no duplication)
	counts := make(map[string]set)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Must set file(s) as argument.\n")
		os.Exit(1)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "File open failed: %v\n", err)
			}
			storeLines(f, counts)
			f.Close()
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			for k := range files {
				fmt.Printf("@%s\n", k)
			}
			fmt.Printf("%s\n\n", line)
		}
	}
}

func storeLines(f *os.File, counts map[string]set) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// make new set if set does not exist
		if counts[input.Text()] == nil {
			counts[input.Text()] = make(set)
		}
		// set temporal value
		counts[input.Text()][f.Name()] = true
	}
}
