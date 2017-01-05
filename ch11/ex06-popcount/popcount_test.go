package popcount

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(255)
	}
}

func BenchmarkShiftPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShiftPopCount(255)
	}
}

func BenchmarkPopCountClearMinBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClearMinBit(255)
	}
}

func TestPopCount(t *testing.T) {
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
			t.Errorf("ShiftPopCount(%x) Expected:%d Actual:%d", test.n, test.want, actual)
		}
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

func TestPopCountClearMinBit(t *testing.T) {
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
		actual := PopCountClearMinBit(test.n)
		if actual != test.want {
			t.Errorf("PopCountClearMinBit(%x) Expected:%d Actual:%d", test.n, test.want, actual)
		}
	}
}
