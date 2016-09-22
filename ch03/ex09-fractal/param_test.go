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
				"x":     {"1.5"},
				"y":     {"2.0"},
				"scale": {"3.0"},
			},
			param{
				x:     1.5,
				y:     2.0,
				scale: 3.0,
			},
		}, {
			url.Values{},
			param{
				x:     0.0,
				y:     0.0,
				scale: 1.0,
			},
		}, {
			url.Values{
				"scale": {"0.0"},
			},
			param{
				x:     0.0,
				y:     0.0,
				scale: 1.0,
			},
		},
	}

	for _, test := range tests {
		got := mapToParam(test.f)
		if got != test.want {
			t.Errorf("Actual: %v\tExpected: %v", got, test.want)
		}
	}
}
