package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    []int
		want bool
	}{
		{
			[]int{1, 3, 5},
			false,
		},
		{
			[]int{1, 3, 1},
			true,
		},
		{
			[]int{1, 3, 3, 1},
			true,
		},
		{
			[]int{1},
			true,
		},
		{
			[]int{},
			true,
		},
	}

	for _, test := range tests {
		got := IsPalindrome(sort.IntSlice(test.s))
		if got != test.want {
			t.Errorf("expected:%v actual:%v", test.want, got)
		}
	}
}
