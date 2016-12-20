package popcount

import (
	"testing"
)

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
