package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/ysohta/gopl-ex/ch03/ex08-newtonsmethod/cmplx128"
	"github.com/ysohta/gopl-ex/ch03/ex08-newtonsmethod/cmplx64"
	"github.com/ysohta/gopl-ex/ch03/ex08-newtonsmethod/cmplxbgflt"
	"github.com/ysohta/gopl-ex/ch03/ex08-newtonsmethod/cmplxrt"
)

var fGetColor = cmplx128.GetColor
var path = "out.png"

// Example of usage
// ./ex08-newtonsmethod cmplx128 cmplx128.png
// ./ex08-newtonsmethod cmplx64 cmplx64.png
// ./ex08-newtonsmethod cmplxbgflt cmplxbgflt.png
// ./ex08-newtonsmethod cmplxrt cmplxrt.png
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	setup(os.Args)

	now := time.Now()

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		elapsed := time.Now().Sub(now)
		fmt.Printf("%4d/%4d\ttime=%dmsec\n", py, height, int(elapsed.Nanoseconds()/1000000))
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			img.Set(px, py, fGetColor(x, y))
		}
	}

	f, _ := os.Create(path)

	png.Encode(f, img)
}

func setup(args []string) {
	if len(args) < 2 {
		return
	}

	switch args[1] {
	case "cmplx64":
		fGetColor = cmplx64.GetColor
	case "cmplxbgflt":
		fGetColor = cmplxbgflt.GetColor
	case "cmplxrt":
		fGetColor = cmplxrt.GetColor
	default:
		fGetColor = cmplx128.GetColor
	}

	if len(args) < 3 {
		return
	}

	path = args[2]
}
