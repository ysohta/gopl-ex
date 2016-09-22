package cmplxbgflt

import (
	"math/big"
	"testing"
)

var two = new(big.Float).SetPrec(prec).SetInt64(2)
var minustwo = new(big.Float).SetPrec(prec).SetInt64(-2)
var five = new(big.Float).SetPrec(prec).SetInt64(5)

func TestAdd(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		y    *complexFloat
		want *complexFloat
	}{
		{
			&complexFloat{one, one},
			&complexFloat{zero, one},
			&complexFloat{one, two},
		}, {
			&complexFloat{minusone, minusone},
			&complexFloat{one, two},
			&complexFloat{zero, one},
		},
	}

	for _, test := range tests {
		target := add(test.x, test.y)
		if target.r.Cmp(test.want.r) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target.r, test.want.r)
		}
		if target.i.Cmp(test.want.i) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target.i, test.want.i)
		}
	}
}

func TestNeg(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		want *complexFloat
	}{
		{
			&complexFloat{one, minusone},
			&complexFloat{minusone, one},
		},
	}

	for _, test := range tests {
		target := neg(test.x)
		if target.r.Cmp(test.want.r) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target, test.want)
		}
		if target.i.Cmp(test.want.i) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target.i, test.want.i)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		y    *complexFloat
		want *complexFloat
	}{
		{
			&complexFloat{one, one},
			&complexFloat{one, minusone},
			&complexFloat{two, zero},
		}, {
			&complexFloat{half, half},
			&complexFloat{half, half},
			&complexFloat{zero, half},
		},
	}

	for _, test := range tests {
		target := mul(test.x, test.y)
		if target.r.Cmp(test.want.r) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target.r, test.want.r)
		}
		if target.i.Cmp(test.want.i) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target.i, test.want.i)
		}
	}
}

func TestQuo(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		y    *complexFloat
		want *complexFloat
	}{
		{
			&complexFloat{one, one},
			&complexFloat{one, minusone},
			&complexFloat{zero, one},
		},
		{
			&complexFloat{two, one},
			&complexFloat{one, minustwo},
			&complexFloat{zero, one},
		},
	}

	for _, test := range tests {
		target := quo(test.x, test.y)
		if target.r.Cmp(test.want.r) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target, test.want)
		}
		if target.i.Cmp(test.want.i) != 0 {
			t.Errorf("Actual(%v)\tExpected(%v)", target.i, test.want.i)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		x    *complexFloat
		want float64
	}{
		{
			&complexFloat{three, four},
			5,
		},
	}

	for _, test := range tests {
		target := abs(test.x)
		if i, _ := target.Float64(); i != test.want {
			t.Errorf("Actual(%v)\tExpected(%v)", target, test.want)
		}
	}
}
