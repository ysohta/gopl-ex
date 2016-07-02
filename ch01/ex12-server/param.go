package main

import (
	"net/url"
	"strconv"
)

type param struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func newParam() param {
	return param{5, 0.001, 100, 64, 8}
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
		case "cycles":
			p.cycles = parseInt(first, 5)
		case "res":
			p.res = parseFloat64(first, 0.001)
		case "size":
			p.size = parseInt(first, 100)
		case "nframes":
			p.nframes = parseInt(first, 64)
		case "delay":
			p.delay = parseInt(first, 8)
		}
	}

	return p
}

func parseInt(str string, defaultVal int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return defaultVal
	}
	return i
}

func parseFloat64(str string, defaultVal float64) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return defaultVal
	}
	return f
}

func compare(a, b param) bool {
	if &a == &b {
		return true
	}

	if a.cycles != b.cycles {
		return false
	}
	if a.res != b.res {
		return false
	}
	if a.size != b.size {
		return false
	}
	if a.nframes != b.nframes {
		return false
	}
	if a.delay != b.delay {
		return false
	}
	return true
}
