package main

import (
	"image/color"
	"net/url"
	"testing"
)

var (
	red  = color.RGBA{0xff, 0x00, 0x00, 0xff}
	blue = color.RGBA{0x00, 0x00, 0xff, 0xff}
)

func TestMapToParam(t *testing.T) {
	tests := []struct {
		f    url.Values
		want param
	}{
		{
			url.Values{
				"width": {"20"},
				"other": {"other"},
			},
			param{
				width:       20,
				height:      320,
				topColor:    red,
				bottomColor: blue,
			},
		}, {
			url.Values{
				"width":       {"400"},
				"height":      {"200"},
				"topColor":    {"0000ffff"},
				"bottomColor": {"ff0000ff"},
			},
			param{
				width:       400,
				height:      200,
				topColor:    blue,
				bottomColor: red,
			},
		},
	}

	for _, test := range tests {
		got := mapToParam(test.f)
		if !compare(got, test.want) {
			t.Errorf("Actual: %v\tExpected: %v", got, test.want)
		}
	}
}

func TestCompare(t *testing.T) {
	p := newParam()
	tests := []struct {
		a    param
		b    param
		want bool
	}{
		{
			p, p, true,
		},
		{
			p, param{}, false,
		},
		{
			p, param{width: 5}, false,
		},
		{
			p,
			param{
				width:  600,
				height: 3,
			},
			false,
		},
		{
			p,
			param{
				width:    600,
				height:   320,
				topColor: blue,
			},
			false,
		},
		{
			p,
			param{
				width:       600,
				height:      320,
				topColor:    red,
				bottomColor: red,
			},
			false,
		},
		{
			p,
			param{
				width:       600,
				height:      320,
				topColor:    red,
				bottomColor: blue,
			},
			true,
		},
	}
	for _, test := range tests {
		got := compare(test.a, test.b)
		if got != test.want {
			t.Errorf("Actual: %v\tExpected: %v", got, test.want)
		}
	}
}
