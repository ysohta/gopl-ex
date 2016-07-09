package main

import (
	"image/color"
	"io"
	"io/ioutil"
	"testing"
)

func TestSurface(t *testing.T) {
	tests := []struct {
		out  io.Writer
		p    param
		want struct {
			width, height         float64
			xyscale, zscale       float64
			topColor, bottomColor color.RGBA
		}
	}{
		{ioutil.Discard,
			param{
				200, 300,
				color.RGBA{0xff, 0xff, 0xff, 0xff},
				color.RGBA{0x0, 0x00, 0x00, 0xff},
			},
			struct {
				width, height         float64
				xyscale, zscale       float64
				topColor, bottomColor color.RGBA
			}{
				200, 300,
				5, 120,
				color.RGBA{0xff, 0xff, 0xff, 0xff},
				color.RGBA{0x0, 0x00, 0x00, 0xff},
			},
		},
	}

	for _, test := range tests {
		surface(test.out, test.p)
		if width != test.want.width &&
			height != test.want.height &&
			xyscale != test.want.xyscale &&
			zscale != test.want.zscale &&
			topColor != test.want.topColor &&
			bottomColor != test.want.bottomColor {
			t.Errorf("Actual(width=%f,height=%f,xyscale=%f,zscale=%f,topColor=%v,bottomColor=%v)\n",
				width, height, xyscale, zscale, topColor, bottomColor)
			t.Errorf("Expected(width=%f,height=%f,xyscale=%f,zscale=%f,topColor=%v,bottomColor=%v)\n",
				test.want.width, test.want.height, test.want.xyscale, test.want.zscale, test.want.topColor, test.want.bottomColor)
		}
	}

}

func TestCorner(t *testing.T) {
	tests := []struct {
		i  int
		j  int
		ok bool
	}{
		{0, 0, true},
		{50, 50, false},
		{100, 100, true},
	}

	for _, test := range tests {
		_, _, _, ok := corner(test.i, test.j)
		if ok != test.ok {
			t.Errorf("Actual(ok=%t)\n", ok)
			t.Errorf("Expected(ok=%t)\n", test.ok)
		}
	}
}

func TestGetFillColor(t *testing.T) {
	initialize(newParam())
	tests := []struct {
		z float64
		c color.RGBA
	}{
		{-128.0, color.RGBA{0x00, 0x00, 0xff, 0xff}},
		{0.0, color.RGBA{0x7f, 0x00, 0x7f, 0xff}},
		{128.0, color.RGBA{0xff, 0x00, 0x00, 0xff}},
		{-129.0, color.RGBA{0x00, 0x00, 0xff, 0xff}},
		{129.0, color.RGBA{0xff, 0x00, 0x00, 0xff}},
	}

	for _, test := range tests {
		got := getFillColor(test.z)
		if got != test.c {
			t.Errorf("Actual(c=%v)\n", got)
			t.Errorf("Expected(c=%v)\n", test.c)
		}
	}
}
