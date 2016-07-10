package main

import (
	"image/color"
	"testing"
)

func TestIntermediate(t *testing.T) {
	tests := []struct {
		colors []color.Color
		want   color.Color
	}{
		{
			[]color.Color{
				color.RGBA{0xff, 0xff, 0x00, 0x00},
				color.RGBA{0x00, 0xff, 0xff, 0x00},
				color.RGBA{0x00, 0x00, 0xff, 0xff},
				color.RGBA{0xff, 0x00, 0x00, 0xff},
			},
			color.RGBA64{0x7fff, 0x7fff, 0x7fff, 0x7fff},
		}, {
			[]color.Color{
				color.RGBA{},
				color.RGBA{},
				color.RGBA{},
				color.RGBA{},
			},
			color.RGBA{},
		},
	}

	for _, test := range tests {
		ar, ag, ab, aa := intermediate(test.colors).RGBA()
		er, eg, eb, ea := test.want.RGBA()
		if ar != er && ag != eg && ab != eb && aa != ea {
			t.Errorf("Actual: %x%x%x%x\tExpected: %x%x%x%x", ar, ag, ab, aa, er, eg, eb, ea)
		}
	}
}
