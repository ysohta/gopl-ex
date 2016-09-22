package cmplx128

import (
	"image/color"
	"math"
	"math/cmplx"
)

const (
	accuracy = 0.01
	max      = 100
)

var solutions = [...]complex128{1, -1, 1i, -1i}

func GetColor(r float64, i float64) color.Color {
	z := complex(r, i)
	x, rep := approximate(z)
	a := 255 - uint8(rep)*16
	d0 := cmplx.Abs(solutions[0] - x)
	d1 := cmplx.Abs(solutions[1] - x)
	d2 := cmplx.Abs(solutions[2] - x)
	d3 := cmplx.Abs(solutions[3] - x)
	min := math.Min(d0, math.Min(d1, math.Min(d2, d3)))
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

func approximate(x complex128) (complex128, int) {
	cnt := 0
	for cmplx.Abs(f(x)) > accuracy && cnt < max {
		x = x - f(x)/fd(x) // approximate value
		cnt++
	}
	return x, cnt
}

func f(x complex128) complex128 {
	return cmplx.Pow(x, 4) - 1
}

func fd(x complex128) complex128 {
	return 4 * cmplx.Pow(x, 3)
}
