package main

import (
	"image/color"
	"testing"
)

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
