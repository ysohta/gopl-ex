package rev

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		arr  [6]int
		want [6]int
	}{
		{
			[...]int{0, 1, 2, 3, 4, 5},
			[...]int{5, 4, 3, 2, 1, 0},
		},
	}

	for _, test := range tests {
		reverse(&test.arr)
		if test.arr != test.want {
			t.Errorf("Expected:%v Actual:%v", test.want, test.arr)
		}
	}
}
