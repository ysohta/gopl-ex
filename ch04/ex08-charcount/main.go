package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

var r = os.Stdin
var w = os.Stdout

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	classification := make(map[string]int)
	invalid := 0

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		for _, c := range unicodeClassification(r) {
			classification[c]++
		}
	}
	fmt.Fprintf(w, "rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Fprint(w, "\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Fprintf(w, "%d\t%d\n", i, n)
		}
	}
	fmt.Fprintf(w, "\nclass\tcount\n")
	for c, n := range classification {
		fmt.Printf("%s\t%d\n", c, n)
	}
}

func unicodeClassification(r rune) []string {
	var s []string

	if unicode.IsControl(r) {
		s = append(s, "Control")
	}
	if unicode.IsDigit(r) {
		s = append(s, "Digit")
	}
	if unicode.IsGraphic(r) {
		s = append(s, "Graphic")
	}
	if unicode.IsLetter(r) {
		s = append(s, "Letter")
	}
	if unicode.IsLower(r) {
		s = append(s, "Lower")
	}
	if unicode.IsMark(r) {
		s = append(s, "Mark")
	}
	if unicode.IsNumber(r) {
		s = append(s, "Number")
	}
	if unicode.IsPrint(r) {
		s = append(s, "Print")
	}
	if unicode.IsPunct(r) {
		s = append(s, "Punct")
	}
	if unicode.IsSpace(r) {
		s = append(s, "Space")
	}
	if unicode.IsSymbol(r) {
		s = append(s, "Symbol")
	}
	if unicode.IsTitle(r) {
		s = append(s, "Title")
	}
	if unicode.IsUpper(r) {
		s = append(s, "Upper")
	}

	return s
}
