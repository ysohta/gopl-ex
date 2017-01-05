package intset

import (
	"bytes"
	"fmt"
)

type IntSetBuiltin map[int]bool

func NewIntSetBuiltin(vals []int) *IntSetBuiltin {
	s := map[int]bool{}
	for _, x := range vals {
		s[x] = true
	}
	isb := IntSetBuiltin(s)
	return &isb
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSetBuiltin) Has(x int) bool {
	return (*s)[x]
}

// Add adds the non-negative value x to the set.
func (s *IntSetBuiltin) Add(x int) {
	(*s)[x] = true
}

// UnionWith sets s to the union of s and t.
func (s *IntSetBuiltin) UnionWith(t *IntSetBuiltin) {
	for k, v := range *t {
		if v {
			(*s)[k] = true
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSetBuiltin) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for k, v := range *s {
		if v {
			if buf.Len() > 1 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(&buf, "%d", k)
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
