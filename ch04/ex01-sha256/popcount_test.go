package main

import (
	"testing"
)

func TestPopcount(t *testing.T) {
	tests := []struct {
		n    []byte
		want int
	}{
		{[]byte{0x00}, 0},
		{[]byte{0x01}, 1},
		{[]byte{0x03}, 2},
		{[]byte{0x10}, 1},
		{[]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, 64},
	}

	for _, test := range tests {
		actual := PopCount(test.n)
		if actual != test.want {
			t.Errorf("PopCount(%x) Expected:%d Actual:%d", test.n, test.want, actual)
		}
	}
}
