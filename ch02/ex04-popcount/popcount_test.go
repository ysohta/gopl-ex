package popcount

import (
	"testing"
)

// => "BenchmarkShiftPopCount-4	30000000	        53.8 ns/op"
func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(0xffffffffffffffff)
	}
}

func TestShiftPopCount(t *testing.T) {
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
		actual := ShiftPopCount(test.n)
		if actual != test.want {
			t.Errorf("ShiftPopCount(%x) Expected:%d Actual:%d", test.n, test.want, actual)
		}
	}
}
