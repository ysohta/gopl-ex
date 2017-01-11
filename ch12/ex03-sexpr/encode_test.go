package sexpr

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
		err  error
	}{
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
			t.Errorf("got:%s want:%s", got, test.want)
		}

		if fmt.Sprint(err) != fmt.Sprint(test.err) {
			t.Errorf("got:%s want:%s", err, test.err)
		}

	}
}
