package intset

import "testing"

func TestHas(t *testing.T) {
	tests := []struct {
		vals []int
	}{
		{[]int{1, 65}},
		{[]int{}},
		{[]int{1000}},
	}

	for _, test := range tests {
		target := NewIntSet(test.vals)
		builtin := NewIntSetBuiltin(test.vals)

		for i := 0; i < 100; i++ {
			got := target.Has(i)
			want := builtin.Has(i)
			if got != want {
				t.Errorf("got:%t want:%t at:%d", got, want, i)
			}
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		vals []int
		x    int
	}{
		{
			[]int{1},
			65,
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.vals)
		builtin := NewIntSetBuiltin(test.vals)

		target.Add(test.x)
		builtin.Add(test.x)

		got := target.String()
		want := target.String()
		if got != want {
			t.Errorf("got:%t want:%t", got, want)
		}
	}
}

func TestUnionwith(t *testing.T) {
	tests := []struct {
		x []int
		y []int
	}{
		{
			[]int{1},
			[]int{65},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{},
			[]int{3},
		},
	}

	for _, test := range tests {
		target := NewIntSet(test.x)
		builtin := NewIntSetBuiltin(test.x)

		target.UnionWith(NewIntSet(test.y))
		builtin.UnionWith(NewIntSetBuiltin(test.y))

		got := target.String()
		want := target.String()
		if got != want {
			t.Errorf("got:%t want:%t", got, want)
		}
	}
}
