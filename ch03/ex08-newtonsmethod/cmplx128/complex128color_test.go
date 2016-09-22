package cmplx128

import (
	"image/color"
	"testing"
)

func TestGetColor(t *testing.T) {
	tests := []struct {
		x    float64
		y    float64
		want color.Color
	}{
		{2, 0, color.RGBA{0xff, 0x00, 0x00, 0xaf}},
		{-2, 0, color.RGBA{0x00, 0xff, 0x00, 0xaf}},
		{0, 2, color.RGBA{0x00, 0x00, 0xff, 0xaf}},
		{0, -2, color.RGBA{0xff, 0xff, 0x00, 0xaf}},
	}

	for _, test := range tests {
		got := GetColor(test.x, test.y)
		if test.want != got {
			t.Errorf("x=%v Actual: %d\tExpected: %d", test.x, got, test.want)
		}
	}
}

func TestApproximate(t *testing.T) {
	tests := []struct {
		x   complex128
		rep int
	}{
		{1, 0},
		{-1, 0},
		{1i, 0},
		{-1i, 0},
		{1.01, 2},
	}

	for _, test := range tests {
		_, got := approximate(test.x)
		if test.rep != got {
			t.Errorf("x=%v Actual: %d\tExpected: %d", test.x, got, test.rep)
		}
	}
}

func TestF(t *testing.T) {
	tests := []struct {
		x complex128
		v complex128
	}{
		{1, 0},
		{2, 15},
	}

	for _, test := range tests {
		got := f(test.x)
		if test.v != got {
			t.Errorf("[v] Actual: %v\tExpected: %v", got, test.v)
		}
	}
}

func TestFd(t *testing.T) {
	tests := []struct {
		x complex128
		v complex128
	}{
		{1, 4},
		{2, 32},
	}

	for _, test := range tests {
		got := fd(test.x)
		if test.v != got {
			t.Errorf("[v] Actual: %v\tExpected: %v", got, test.v)
		}
	}
}
