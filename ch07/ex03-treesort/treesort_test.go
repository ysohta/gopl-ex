package treesort

import "testing"
import (
	"reflect"
)

func TestString(t *testing.T) {
	tests := []struct {
		values []int
		want   string
	}{
		{
			[]int{3, -1, 7, 0},
			"-1 0 3 7",
		},
		{
			[]int{0, 1, 2, 3},
			"0 1 2 3",
		},
		{
			[]int{0, -1, -2, -3},
			"-3 -2 -1 0",
		},
		{
			[]int{1},
			"1",
		},
		{
			[]int{},
			"",
		},
	}

	for _, test := range tests {
		var target *tree
		for _, value := range test.values {
			target = add(target, value)
		}

		got := target.String()
		if got != test.want {
			t.Errorf("expected:%s actual:%s", test.want, got)
		}
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		values []int
		want   []int
	}{
		{
			[]int{3, -1, 7},
			[]int{-1, 3, 7},
		},
	}

	for _, test := range tests {
		Sort(test.values)

		if !reflect.DeepEqual(test.values, test.want) {
			t.Errorf("expected:%v actual:%v", test.want, test.values)
		}
	}
}
