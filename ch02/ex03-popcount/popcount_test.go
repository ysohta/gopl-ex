package popcount

import (
	"testing"
)

// => "BenchmarkPopcount-4    	200000000	         6.87 ns/op"
func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0xffffffffffffffff)
	}
}

// => "BenchmarkLoopPopCount-4	100000000	        17.2 ns/op"
func BenchmarkLoopPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LoopPopCount(0xffffffffffffffff)
	}
}

func TestPopcount(t *testing.T) {
	tests := []struct {
		n    uint64
		want int
	}{
		{0x00, 0},
		{0x01, 1},
		{0x03, 2},
		{0x10, 1},
		{0xffffffffffffffff, 64},
	}

	for _, test := range tests {
		actual := PopCount(test.n)
		if actual != test.want {
			t.Errorf("PopCount(%x) Expected:%d Actual:%d", test.n, test.want, actual)
		}
	}
}

func TestLoopPopCount(t *testing.T) {
	tests := []struct {
		n    uint64
		want int
	}{
		{0x00, 0},
		{0x01, 1},
		{0x03, 2},
		{0x10, 1},
		{0xffffffffffffffff, 64},
	}

	for _, test := range tests {
		actual := LoopPopCount(test.n)
		if actual != test.want {
			t.Errorf("LoopPopCount(%x) Expected:%d Actual:%d", test.n, test.want, actual)
		}
	}
}
