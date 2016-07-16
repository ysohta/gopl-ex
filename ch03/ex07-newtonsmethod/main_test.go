package main

import (
	"image/color"
	"testing"
)

func TestGetColor(t *testing.T) {
	tests := []struct {
		x    complex128
		rep  int
		want color.Color
	}{
		{2, 10, color.RGBA{0xff, 0x00, 0x00, 0x5f}},
		{-2, 10, color.RGBA{0x00, 0xff, 0x00, 0x5f}},
		{2i, 10, color.RGBA{0x00, 0x00, 0xff, 0x5f}},
		{-2i, 10, color.RGBA{0xff, 0xff, 0x00, 0x5f}},
	}

	for _, test := range tests {
		got := getColor(test.x, test.rep)
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
		{1, 3},
		{2, 24},
	}

	for _, test := range tests {
		got := fd(test.x)
		if test.v != got {
			t.Errorf("[v] Actual: %v\tExpected: %v", got, test.v)
		}
	}
}
