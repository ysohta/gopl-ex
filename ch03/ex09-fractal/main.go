package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
)

const (
	width, height = 1024, 1024
	dist          = 2.0
)

func main() {
	// Query 	Example:
	// http://localhost:8000/?x=0&y=1&scale=5
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		p := mapToParam(r.Form)
		fractal(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func fractal(out io.Writer, p param) {
	scaled := dist / p.scale
	xmin := -scaled + p.x
	xmax := scaled + p.x
	ymin := -scaled - p.y
	ymax := scaled - p.y

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, madelbrot(z))
		}
	}
	png.Encode(out, img)
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
