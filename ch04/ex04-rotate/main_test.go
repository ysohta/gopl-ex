package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		s     []int
		shift int
		want  []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5},
			-2,
			[]int{2, 3, 4, 5, 0, 1},
		}, {
			[]int{0, 1, 2, 3, 4, 5},
			3,
			[]int{3, 4, 5, 0, 1, 2},
		}, {
			[]int{0, 1, 2, 3, 4, 5},
			6,
			[]int{0, 1, 2, 3, 4, 5},
		}, {
			[]int{0, 1, 2, 3, 4, 5},
			1,
			[]int{5, 0, 1, 2, 3, 4},
		}, {
			[]int{0, 1, 2, 3, 4, 5},
			7,
			[]int{5, 0, 1, 2, 3, 4},
		},
	}

	for _, test := range tests {
		rotate(test.s, test.shift)
		if !equal(test.s, test.want) {
			t.Errorf("Expected:%v Actual:%v", test.want, test.s)
		}
	}
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
