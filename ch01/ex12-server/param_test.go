package main

import (
	"net/url"
	"testing"
)

func TestMapToParam(t *testing.T) {
	tests := []struct {
		f    url.Values
		want param
	}{
		{
			url.Values{
				"cycles": {"20"},
				"res":    {"0.005"},
				"other":  {"other"},
			},
			param{
				cycles:  20,
				res:     0.005,
				size:    100,
				nframes: 64,
				delay:   8,
			},
		}, {
			url.Values{
				"cycles":  {"20"},
				"res":     {"0.1"},
				"size":    {"10"},
				"nframes": {"20"},
				"delay":   {"30"},
			},
			param{
				cycles:  20,
				res:     0.1,
				size:    10,
				nframes: 20,
				delay:   30,
			},
		}, {
			url.Values{},
			param{
				cycles:  5,
				res:     0.001,
				size:    100,
				nframes: 64,
				delay:   8,
			},
		}, {
			url.Values{
				"cycles":  {"a"},
				"res":     {"b"},
				"size":    {"c"},
				"nframes": {"d", "e"},
				"delay":   {""},
			},
			param{
				cycles:  5,
				res:     0.001,
				size:    100,
				nframes: 64,
				delay:   8,
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
			p, param{cycles: 5}, false,
		},
		{
			p,
			param{
				cycles: 5,
				res:    0.001,
			},
			false,
		},
		{
			p,
			param{
				cycles: 5,
				res:    0.001,
				size:   100,
			},
			false,
		},
		{
			p,
			param{
				cycles:  5,
				res:     0.001,
				size:    100,
				nframes: 64,
			},
			false,
		},
		{
			p,
			param{
				cycles:  5,
				res:     0.001,
				size:    100,
				nframes: 64,
				delay:   8,
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
