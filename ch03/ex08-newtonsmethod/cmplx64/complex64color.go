package cmplx64

import (
	"image/color"
	"math/cmplx"
)

const (
	accuracy = 0.01
	max      = 100
)

var solutions = [...]complex64{1, -1, 1i, -1i}

func GetColor(r float64, i float64) color.Color {
	var z complex64
	z = complex(float32(r), float32(i))
	x, rep := approximate(z)
	a := 255 - uint8(rep)*16
	d0 := abs(solutions[0] - x)
	d1 := abs(solutions[1] - x)
	d2 := abs(solutions[2] - x)
	d3 := abs(solutions[3] - x)

	// find minimum
	min := d0
	if min > d1 {
		min = d1
	}
	if min > d2 {
		min = d2
	}
	if min > d3 {
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

func approximate(x complex64) (complex64, int) {
	cnt := 0
	for abs(f(x)) > accuracy && cnt < max {
		x = x - f(x)/fd(x) // approximate value
		cnt++
	}
	return x, cnt
}

func f(x complex64) complex64 {
	return pow(x, 4) - 1
}

func fd(x complex64) complex64 {
	return 4 * pow(x, 3)
}

func abs(x complex64) float32 {
	return float32(cmplx.Abs(complex128(x)))
}

func pow(x complex64, y complex64) complex64 {
	return complex64(cmplx.Pow(complex128(x), complex128(y)))
}
