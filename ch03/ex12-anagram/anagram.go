// Package strings implements functions to handle string.
package strings

import (
	"unicode"
)

func anagram(s1, s2 string) bool {
	m1 := freq(s1)
	m2 := freq(s2)

	if len(m1) != len(m2) {
		return false
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}

func freq(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		if !unicode.IsSpace(r) {
			m[unicode.ToLower(r)]++
		}
	}
	return m
}
