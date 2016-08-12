package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%s\n", reverse([]byte("Hello, 世界"))) // => "界世 ,olleH"
}

func reverse(s []byte) []byte {
	start, end := 0, len(s)
	length := utf8.RuneCount(s)
	var cnt int
	for cnt < length/2 {
		_, size1 := utf8.DecodeRune(s[start:])
		_, size2 := utf8.DecodeLastRune(s[:end])

		swap(s, start, size1, end-size2, size2)

		start += size2
		end -= size1
		cnt++
	}

	return s
}

// swap returns slice whose elements with specified index and size are swapped.
func swap(s []byte, i1, s1, i2, s2 int) []byte {
	// move forward
	for i := 0; i < s2; i++ {
		move(s, i2+i, i1+i)
	}

	// move backward
	offset := i1 + s2
	diff := i2 - i1 - 1
	for i := 0; i < s1; i++ {
		move(s, offset, offset+diff)
	}

	return s
}

// move returns slice whose element is moved at specified index.
func move(s []byte, from, to int) []byte {
	if from < to {
		for i := from; i < to; i++ {
			s[i], s[i+1] = s[i+1], s[i]
		}
	} else {
		for i := from; i > to; i-- {
			s[i], s[i-1] = s[i-1], s[i]
		}
	}
	return s
}
