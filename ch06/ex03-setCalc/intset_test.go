package intset

import (
	"testing"
)

func TestUnionWith(t *testing.T) {
	tests := []struct {
		x    []int
		y    []int
		want string
	}{
		{
			[]int{1, 144, 9},
			[]int{8, 9, 10},
			"{1 8 9 10 144}",
		}, {
			[]int{1, 144, 9},
			[]int{1},
			"{1 9 144}",
		}, {
			[]int{},
			[]int{1},
			"{1}",
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		y := NewIntSet(test.y)
		x.UnionWith(y)
		got := x.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		x    []int
		y    []int
		want string
	}{
		{
			[]int{1, 144, 9},
			[]int{8, 9, 1},
			"{1 9}",
		}, {
			[]int{1, 144, 9},
			[]int{3, 4, 5},
			"{}",
		}, {
			[]int{1, 144, 9},
			[]int{1},
			"{1}",
		}, {
			[]int{1},
			[]int{1, 144, 9},
			"{1}",
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		y := NewIntSet(test.y)
		x.IntersectWith(y)
		got := x.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		x    []int
		y    []int
		want string
	}{
		{
			[]int{1, 144, 9},
			[]int{8, 9, 1},
			"{144}",
		}, {
			[]int{1, 144, 9},
			[]int{3, 4, 5},
			"{1 9 144}",
		}, {
			[]int{1, 144, 9},
			[]int{1},
			"{9 144}",
		}, {
			[]int{1},
			[]int{1, 144, 9},
			"{}",
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		y := NewIntSet(test.y)
		x.DifferenceWith(y)
		got := x.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		x    []int
		y    []int
		want string
	}{
		{
			[]int{1, 144, 9},
			[]int{8, 9, 1},
			"{8 144}",
		}, {
			[]int{1, 144, 9},
			[]int{3, 4, 5},
			"{1 3 4 5 9 144}",
		}, {
			[]int{1, 144, 9},
			[]int{1},
			"{9 144}",
		}, {
			[]int{1},
			[]int{1, 144, 9},
			"{9 144}",
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		y := NewIntSet(test.y)
		x.SymmetricDifference(y)
		got := x.String()
		if got != test.want {
			t.Errorf("expected=%s actual=%s", test.want, got)
		}
	}
}
