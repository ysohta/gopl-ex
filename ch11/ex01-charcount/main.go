package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var (
	out io.Writer = os.Stdout
)

type result struct {
	counts  map[rune]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func main() {
	res := charcount(os.Stdin)
	printCounts(res)
}

func charcount(r io.Reader) result {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return result{counts, utflen, invalid}
}

func printCounts(res result) {
	fmt.Fprintf(out, "rune\tcount\n")
	for c, n := range res.counts {
		fmt.Fprintf(out, "%q\t%d\n", c, n)
	}

	fmt.Fprintf(out, "\nlen\tcount\n")
	for i, n := range res.utflen {
		if i > 0 {
			fmt.Fprintf(out, "%d\t%d\n", i, n)
		}
	}
	if res.invalid > 0 {
		fmt.Fprintf(out, "\n%d invalid UTF-8 characters\n", res.invalid)
	}
}
