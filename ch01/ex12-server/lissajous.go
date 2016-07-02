// ex06-lissajous creates random lissajous figure with multi-colored curves.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var pallete = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff}}

const (
	blackIndex = 0
	redIndex   = 1
	greenIndex = 2
	blueIndex  = 3
)

func lissajous(out io.Writer, p param) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: p.nframes}
	phase := 0.0
	for i := 0; i < p.nframes; i++ {
		rect := image.Rect(0, 0, 2*p.size+1, 2*p.size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < float64(p.cycles)*2*math.Pi; t += p.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// choose colorIndex
			img.SetColorIndex(p.size+int(x*float64(p.size)+0.5), p.size+int(y*float64(p.size)+0.5), colorIndex(t))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func colorIndex(t float64) uint8 {
	if math.Sin(t) > 0 && math.Cos(t) >= -0.5 {
		return redIndex
	} else if math.Sin(t) <= 0 && math.Cos(t) > -0.5 {
		return greenIndex
	} else {
		return blueIndex
	}
}
