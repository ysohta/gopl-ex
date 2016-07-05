package main

import (
	"testing"
)

func TestCorner(t *testing.T) {
	tests := []struct {
		i  int
		j  int
		ok bool
	}{
		{0, 0, true},
		{50, 50, false},
		{100, 100, true},
	}

	for _, test := range tests {
		_, _, ok := corner(test.i, test.j)
		if ok != test.ok {
			t.Errorf("Actual(ok=%t)\n", ok)
			t.Errorf("Expected(ok=%t)\n", test.ok)
		}
	}
}
