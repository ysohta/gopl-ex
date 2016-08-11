package main

import (
	"testing"
)

func TestDiffBitSha256(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{
			"hoge",
			"hoge",
			0,
		}, {
			"x",
			"X",
			125,
		},
	}

	for _, test := range tests {
		got := diffBitSha256(test.s1, test.s2)
		if got != test.want {
			t.Errorf("Expected:%d Actual:%d", test.want, got)
		}
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		b1   [32]byte
		b2   [32]byte
		want [32]byte
	}{
		{
			[32]byte{1, 2, 3, 4, 5, 6, 7, 8},
			[32]byte{1, 2, 3, 4, 5, 6, 7, 8},
			[32]byte{},
		}, {
			[32]byte{0xff, 0xff, 0xff, 0xff},
			[32]byte{},
			[32]byte{0xff, 0xff, 0xff, 0xff},
		},
	}

	for _, test := range tests {
		got := xor(test.b1, test.b2)
		if got != test.want {
			t.Errorf("Expected:%x Actual:%x", test.want, got)
		}
	}
}
