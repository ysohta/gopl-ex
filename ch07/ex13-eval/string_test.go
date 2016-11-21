package eval

import (
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}},
		{"5 / 9 * (F - 32)", Env{"F": -40}},
		{"5 / 9 * (F - 32)", Env{"F": 32}},
		{"5 / 9 * (F - 32)", Env{"F": 212}},
		{"-1 + -x", Env{"x": 1}},
		{"-1 - x", Env{"x": 1}},
		{"(x - 1) * (x + 1)", Env{"x": -1}},
	}
	for _, test := range tests {
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		got := expr.String()
		if got != test.expr {
			t.Errorf("got %q want %q\n", got, test.expr)
		}
	}
}
