// ex03-surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var topColor = color.RGBA{0xff, 0x00, 0x00, 0xff}
var bottomColor = color.RGBA{0x00, 0x00, 0xff, 0xff}

func main() {
	fmt.Printf("<svg xmlns ='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			if ax, ay, az, ok := corner(i+1, j); !ok {
				continue
			} else if bx, by, bz, ok := corner(i, j); !ok {
				continue
			} else if cx, cy, cz, ok := corner(i, j+1); !ok {
				continue
			} else if dx, dy, dz, ok := corner(i+1, j+1); !ok {
				continue
			} else {
				mean := (az + bz + cz + dz) / 4
				c := getFillColor(mean)
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%02x%02x%02x'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, c.R, c.G, c.B)
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64, bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	sz := z * zscale
	return sx, sy, sz, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func getFillColor(h float64) color.RGBA {
	scale := (h + zscale) / (zscale * 2)
	if scale < 0.0 {
		scale = 0.0
	} else if scale > 1.0 {
		scale = 1.0
	}

	r := uint8((float64(topColor.R)-float64(bottomColor.R))*scale + float64(bottomColor.R))
	g := uint8((float64(topColor.G)-float64(bottomColor.G))*scale + float64(bottomColor.G))
	b := uint8((float64(topColor.B)-float64(bottomColor.B))*scale + float64(bottomColor.B))
	a := uint8((float64(topColor.A)-float64(bottomColor.A))*scale + float64(bottomColor.A))

	return color.RGBA{r, g, b, a}
}
