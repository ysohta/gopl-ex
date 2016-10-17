package intset

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		data []int
		vals []int
		want string
	}{
		{
			[]int{1, 144, 9},
			[]int{3, 5},
			"{1 3 5 9 144}",
		}, {
			[]int{1, 144, 9},
			[]int{1, 9},
			"{1 9 144}",
		}, {
			[]int{},
			[]int{1, 9},
			"{1 9}",
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.data)
		target.AddAll(test.vals...)
		got := target.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}
