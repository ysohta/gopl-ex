package intset

import (
	"reflect"
	"testing"
)

func TestElems(t *testing.T) {
	tests := []struct {
		x    []int
		want []int
	}{
		{
			[]int{1, 144, 9},
			[]int{1, 9, 144},
		}, {
			[]int{1, 9},
			[]int{1, 9},
		}, {
			[]int{},
			[]int{},
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		got := x.Elems()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("expected=%v actual=%v", test.want, got)
		}
	}
}
