package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
)

var (
	num int
	out io.Writer = os.Stdout
)

func init() {
	flag.IntVar(&num, "n", runtime.NumCPU(), "number of goroutine")
	flag.Parse()
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	var wg sync.WaitGroup

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < num; i++ {
		from := height / num * i
		to := height / num * (i + 1)

		wg.Add(1)
		go func(from, to int) {
			for py := from; py < to; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					// Image point (px, py) represents complex value z.
					img.Set(px, py, mandelbrot(z))
				}
			}
			wg.Done()
		}(from, to)
	}
	wg.Wait()

	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
