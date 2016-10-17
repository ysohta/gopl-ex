package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func NewIntSet(vals []int) *IntSet {
	intSet := &IntSet{}

	for _, v := range vals {
		intSet.Add(v)
	}

	return intSet
}

func (s *IntSet) Len() int {
	return len(s.words)
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if s.Has(x) {
		s.words[word] ^= (1 << bit)
	}
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	copied := &IntSet{make([]uint64, len(s.words))}
	copy(copied.words, s.words)
	return copied
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
