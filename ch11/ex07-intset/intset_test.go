package intset

import (
	"math/rand"
	"testing"
	"time"
)

var (
	v1 int
	v2 int
)

func init() {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	v1 = rng.Intn(0xFFFF)
	v2 = rng.Intn(0xFFFF)
}

func BenchmarkIntSetHas(b *testing.B) {
	target := NewIntSet([]int{v1})
	for i := 0; i < b.N; i++ {
		target.Has(v2)
	}
}

func BenchmarkIntSet32Has(b *testing.B) {
	target := NewIntSet32([]int{v1})
	for i := 0; i < b.N; i++ {
		target.Has(v2)
	}
}

func BenchmarkIntSetBuiltinHas(b *testing.B) {
	target := NewIntSetBuiltin([]int{v1})
	for i := 0; i < b.N; i++ {
		target.Has(v2)
	}
}

func BenchmarkIntAdd(b *testing.B) {
	target := NewIntSet([]int{v1})
	for i := 0; i < b.N; i++ {
		target.Add(v2)
	}
}

func BenchmarkIntAdd32(b *testing.B) {
	target := NewIntSet32([]int{v1})
	for i := 0; i < b.N; i++ {
		target.Add(v2)
	}
}

func BenchmarkIntSetBuiltinAdd(b *testing.B) {
	target := NewIntSetBuiltin([]int{v1})
	for i := 0; i < b.N; i++ {
		target.Add(v2)
	}
}

func BenchmarkIntUnionWith(b *testing.B) {
	target := NewIntSet([]int{v1})
	for i := 0; i < b.N; i++ {
		target.UnionWith(NewIntSet([]int{v2}))
	}
}
func BenchmarkIntUnionWith32(b *testing.B) {
	target := NewIntSet32([]int{v1})
	for i := 0; i < b.N; i++ {
		target.UnionWith(NewIntSet32([]int{v2}))
	}
}

func BenchmarkIntSetBuiltinUnionWith(b *testing.B) {
	target := NewIntSetBuiltin([]int{v1})
	for i := 0; i < b.N; i++ {
		target.UnionWith(NewIntSetBuiltin([]int{v2}))
	}
}

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
