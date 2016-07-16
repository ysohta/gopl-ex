package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

const (
	accuracy   = 0.01
	nsolutions = 4
)

var solutions = [...]complex128{1, -1, 1i, -1i}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, getColor(approximate(z)))
		}
	}
	png.Encode(os.Stdout, img)
}

func getColor(x complex128, rep int) color.Color {
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
	for cmplx.Abs(f(x)) > accuracy {
		x = x - f(x)/fd(x) // approximate value
		cnt++
	}
	return x, cnt
}

func f(x complex128) complex128 {
	return cmplx.Pow(x, 4) - 1
}

func fd(x complex128) complex128 {
	return 3 * cmplx.Pow(x, 3)
}
