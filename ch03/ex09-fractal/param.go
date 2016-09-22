package main

import (
	"net/url"
	"strconv"
)

type param struct {
	x     float64
	y     float64
	scale float64
}

func newParam() param {
	return param{0, 0, 1}
}

func mapToParam(m url.Values) param {
	p := newParam()
	for k, v := range m {
		// ignore nil values
		if v == nil {
			continue
		}

		first := v[0]
		switch k {
		case "x":
			p.x = parseFloat64(first, 0)
		case "y":
			p.y = parseFloat64(first, 0)
		case "scale":
			f := parseFloat64(first, 1.0)
			if f > 0 {
				p.scale = f
			}
		}
	}

	return p
}

func parseFloat64(str string, defaultVal float64) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defaultVal
	}
	return f
}
