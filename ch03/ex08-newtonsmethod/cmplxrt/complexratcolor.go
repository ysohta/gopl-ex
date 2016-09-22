package cmplxrt

import (
	"image/color"
	"math/big"
)

const (
	accuracy = 0.01
	max      = 1
)

var zero = new(big.Rat).SetInt64(0)
var one = new(big.Rat).SetInt64(1)
var three = new(big.Rat).SetInt64(3)
var four = new(big.Rat).SetInt64(4)
var minusone = new(big.Rat).SetInt64(-1)
var half = new(big.Rat).SetFloat64(0.5)

var fAccuracy = new(big.Rat).SetFloat64(accuracy)

var solutions = [...]*complexRat{
	&complexRat{one, zero},
	&complexRat{minusone, zero},
	&complexRat{zero, one},
	&complexRat{zero, minusone},
}

func GetColor(r float64, i float64) color.Color {
	z := complexRat{
		new(big.Rat).SetFloat64(r),
		new(big.Rat).SetFloat64(i),
	}
	x, rep := approximate(&z)

	a := 255 - uint8(rep)*16
	d0 := abs(add(solutions[0], neg(x)))
	d1 := abs(add(solutions[1], neg(x)))
	d2 := abs(add(solutions[2], neg(x)))
	d3 := abs(add(solutions[3], neg(x)))

	// find minimum
	min := d0
	if min.Cmp(d1) > 0 {
		min = d1
	}
	if min.Cmp(d2) > 0 {
		min = d2
	}
	if min.Cmp(d3) > 0 {
		min = d3
	}

	switch min {
	case d0:
		return color.RGBA{255, 0, 0, a}
	case d1:
		return color.RGBA{0, 255, 0, a}
	case d2:
		return color.RGBA{0, 0, 255, a}
	case d3:
		return color.RGBA{255, 255, 0, a}
	}
	return color.Black
}

func approximate(x *complexRat) (*complexRat, int) {
	cnt := 0
	for abs(f(x)).Cmp(fAccuracy) > 0 && cnt < max {
		x = add(x, neg(quo(f(x), fd(x))))
		cnt++
	}
	return x, cnt
}

func f(x *complexRat) *complexRat {
	// x^4 - 1
	return add(mul(mul(x, x), mul(x, x)), &complexRat{minusone, zero})
}

func fd(x *complexRat) *complexRat {
	// 4 * x^3
	y := mul(x, mul(x, x))
	y.r.Mul(y.r, four)
	y.i.Mul(y.i, four)
	return y
}
