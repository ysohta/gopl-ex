package sexpr

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestIsZeroValue(t *testing.T) {
	tests := []struct {
		v    interface{}
		want bool
	}{
		{interface{}(0), true},
		{interface{}(-1), false},
		{interface{}(""), true},
		{interface{}("str"), false},
		{interface{}(true), false},
		{interface{}(false), true},
		{interface{}(math.Pi), false},
		{interface{}(0.0), true},
		{interface{}(0 + 0i), true},
		{interface{}(1 + 0i), false},
		{interface{}(-1i), false},
	}

	for _, test := range tests {
		got := isZeroValue(reflect.ValueOf(test.v))
		if got != test.want {
			t.Errorf("%v got:%t want:%t", test.v, got, test.want)
		}
	}
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
		err  error
	}{
		{
			struct {
				x  int
				y  int
				s1 string
				s2 string
			}{x: 3, s2: "str"},
			`((x 3) (s2 "str"))`,
			nil,
		},
		{
			struct {
				x  uint
				y  uint
				b1 bool
				b2 bool
			}{y: 3, b1: true},
			`((y 3) (b1 t))`,
			nil,
		},
		{
			interface{}(true),
			"t",
			nil,
		},
		{
			interface{}(false),
			"nil",
			nil,
		},
		{
			interface{}(3.14),
			"3.14",
			nil,
		},
		{
			interface{}(float32(0.25)),
			"0.25",
			nil,
		},
		{
			interface{}(1 + 2i),
			"#C(1.0 2.0)",
			nil,
		},
		{
			struct {
				x int
				y int
			}{3, 4},
			`((x 3) (y 4))`,
			nil,
		},
		{
			struct{ x interface{} }{[]int{1, 2, 3}},
			`((x ("[]int" (1 2 3))))`,
			nil,
		},
	}

	for _, test := range tests {
		got, err := Marshal(test.v)

		if string(got) != test.want {
			t.Errorf("%v got:%s want:%s", test.v, got, test.want)
		}

		if fmt.Sprint(err) != fmt.Sprint(test.err) {
			t.Errorf("got:%s want:%s", err, test.err)
		}

	}
}
