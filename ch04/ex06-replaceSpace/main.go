package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := "a      b\tc"
	b := []byte(s)
	b = replaceSpace(b)
	fmt.Printf("%s\n", b)
	// => "[a b c]"
}

func replaceSpace(b []byte) []byte {
	out := b[:0]
	length := len(b)
	offset := 0
	for offset < length {
		r, size := utf8.DecodeRune(b)

		if unicode.IsSpace(r) {
			if out[len(out)-1] != ' ' {
				out = append(out, ' ')
			}
		} else {
			out = appendAll(out, b[:size])
		}

		b = b[size:]
		offset += size
	}
	return out
}

func appendAll(s1 []byte, s2 []byte) []byte {
	for _, b := range s2 {
		s1 = append(s1, b)
	}
	return s1
}
