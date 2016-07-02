package main

import (
	"math"
	"testing"
)

func TestColorIndex(t *testing.T) {
	var tests = []struct {
		input float64
		want  uint8
	}{
		{0.0, greenIndex},
		{math.Pi * 1.0 / 2.0, redIndex},
		{math.Pi, blueIndex},
		{math.Pi * 3.0 / 2.0, greenIndex},
		{math.Pi * 2, greenIndex},
	}
	for _, test := range tests {
		if got := colorIndex(test.input); got != test.want {
			t.Errorf("colorIndex(%f) = %d,", test.input, got)
		}
	}
}
