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

func TestString(t *testing.T) {
	tests := []struct {
		x    []int
		want string
	}{
		{
			[]int{1, 144, 9},
			"{1 9 144}",
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		got := x.String()
		if got != test.want {
			t.Errorf("expected=%v actual=%v", test.want, got)
		}
	}
}

func TestWordLen(t *testing.T) {

	tests := []struct {
		x      []int
		want32 int
		want64 int
	}{
		{
			[]int{31},
			1,
			1,
		}, {
			[]int{32},
			2,
			1,
		}, {
			[]int{63},
			2,
			1,
		}, {
			[]int{64},
			3,
			2,
		},
	}

	for _, test := range tests {
		x := NewIntSet(test.x)
		got := len(x.words)
		var want int
		switch size {
		case 32:
			want = test.want32
		case 64:
			want = test.want64
		default:
			t.Errorf("unexpected size:%d", size)
		}

		if got != want {
			t.Errorf("expected=%v actual=%v", want, got)
		}
	}
}
