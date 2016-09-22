package cmplxrt

import (
	"math/big"
	"testing"
)

var two = new(big.Rat).SetInt64(2)
var minustwo = new(big.Rat).SetInt64(-2)
var five = new(big.Rat).SetInt64(5)

func TestAdd(t *testing.T) {
	tests := []struct {
		x    *complexRat
		y    *complexRat
		want *complexRat
	}{
		{
			&complexRat{one, one},
			&complexRat{zero, one},
			&complexRat{one, two},
		}, {
			&complexRat{minusone, minusone},
			&complexRat{one, two},
			&complexRat{zero, one},
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
		x    *complexRat
		want *complexRat
	}{
		{
			&complexRat{one, minusone},
			&complexRat{minusone, one},
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
		x    *complexRat
		y    *complexRat
		want *complexRat
	}{
		{
			&complexRat{one, one},
			&complexRat{one, minusone},
			&complexRat{two, zero},
		}, {
			&complexRat{half, half},
			&complexRat{half, half},
			&complexRat{zero, half},
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
		x    *complexRat
		y    *complexRat
		want *complexRat
	}{
		{
			&complexRat{one, one},
			&complexRat{one, minusone},
			&complexRat{zero, one},
		},
		{
			&complexRat{two, one},
			&complexRat{one, minustwo},
			&complexRat{zero, one},
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
		x     *complexRat
		want  float64
		delta float64
	}{
		{
			&complexRat{three, four},
			5,
			1,
		},
	}

	for _, test := range tests {
		target, _ := abs(test.x).Float64()
		if target < test.want-test.delta || target > test.want+test.delta {
			t.Errorf("Actual(%v)\tExpected(%v)", target, test.want)
		}
	}
}
