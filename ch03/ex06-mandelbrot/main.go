package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
	nsubpix                = 2
	shift                  = 1.0 / nsubpix
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img.Set(px, py, supersampling(float64(px), float64(py)))
		}
	}
	png.Encode(os.Stdout, img)
}

func supersampling(px, py float64) color.Color {
	var pixels []color.Color
	for offsety := 0.0; offsety < 1.0; offsety += shift {
		for offsetx := 0.0; offsetx < 1.0; offsetx += shift {
			y := (py+offsety)/height*(ymax-ymin) + ymin
			x := (px+offsetx)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			pixels = append(pixels, madelbrot(z))
		}
	}
	return intermediate(pixels)
}

func madelbrot(z complex128) color.Color {
	const interations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < interations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{255 - contrast*n, 0, 255}
		}
	}
	return color.Black
}

func intermediate(colors []color.Color) color.Color {
	var sR, sG, sB, sA uint32
	n := uint32(len(colors))
	for _, c := range colors {
		r, g, b, a := c.RGBA()
		sR += r
		sG += g
		sB += b
		sA += a
	}
	return color.RGBA64{
		uint16(sR / n),
		uint16(sG / n),
		uint16(sB / n),
		uint16(sA / n),
	}
}
