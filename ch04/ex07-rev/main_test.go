package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"abc",
			"cba",
		}, {
			"世界",
			"界世",
		}, {
			"魔貫光殺砲",
			"砲殺光貫魔",
		}, {
			"Hello, 世界",
			"界世 ,olleH",
		},
	}

	for _, test := range tests {
		b := []byte(test.s)
		b = reverse(b)
		got := string(b)
		if got != test.want {
			t.Errorf("Expected:%s Actual:%s", test.want, got)
		}
	}
}

func TestSwap(t *testing.T) {
	tests := []struct {
		s              string
		i1, s1, i2, s2 int
		want           string
	}{
		{
			"abc", 0, 1, 2, 1,
			"cba",
		}, {
			"abc", 0, 1, 1, 1,
			"bac",
		}, {
			"abcd", 0, 2, 2, 2,
			"cdab",
		}, {
			"abcde", 0, 1, 3, 2,
			"debca",
		}, {
			"abcde", 0, 3, 3, 2,
			"deabc",
		}, {
			"abcdefghi", 0, 3, 6, 3,
			"ghidefabc",
		},
	}

	for _, test := range tests {
		b := []byte(test.s)
		b = swap(b, test.i1, test.s1, test.i2, test.s2)
		got := string(b)
		if got != test.want {
			t.Errorf("Expected:%s Actual:%s", test.want, got)
		}
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		s        string
		from, to int
		want     string
	}{
		{
			"abc", 0, 2,
			"bca",
		}, {
			"abc", 2, 0,
			"cab",
		},
	}

	for _, test := range tests {
		b := []byte(test.s)
		b = move(b, test.from, test.to)
		got := string(b)
		if got != test.want {
			t.Errorf("Expected:%s Actual:%s", test.want, got)
		}
	}
}
