package intset

import (
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{
			[]int{1, 144, 9},
			3,
		}, {
			[]int{},
			0,
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.data)
		got := target.Len()
		if got != test.want {
			t.Errorf("expected=%d actual=%d", test.want, got)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		data []int
		x    int
		want string
	}{
		{
			[]int{1, 144, 9},
			9,
			"{1 144}",
		}, {
			[]int{1, 144, 9},
			3,
			"{1 9 144}",
		}, {
			[]int{},
			3,
			"{}",
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.data)
		target.Remove(test.x)
		got := target.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}

func TestClear(t *testing.T) {
	tests := []struct {
		data []int
		want string
	}{
		{
			[]int{1, 144, 9},
			"{}",
		}, {
			[]int{},
			"{}",
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.data)
		target.Clear()
		got := target.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}
func TestCopy(t *testing.T) {
	tests := []struct {
		data []int
		want string
	}{
		{
			[]int{1, 144, 9},
			"{1 9 144}",
		}, {
			[]int{},
			"{}",
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.data)
		cp := target.Copy()
		target.Add(20) // add a value
		got := cp.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}
