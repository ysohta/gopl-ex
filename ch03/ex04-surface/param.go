package main

import (
	"image/color"
	"net/url"
	"strconv"
)

type param struct {
	width       int
	height      int
	topColor    color.RGBA
	bottomColor color.RGBA
}

func newParam() param {
	return param{
		600,
		320,
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0x00, 0xff, 0xff},
	}
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
		case "width":
			p.width = parseInt(first, 600)
		case "height":
			p.height = parseInt(first, 320)
		case "topColor":
			p.topColor = getRGBAColor(parseIntHex(first, 0xff0000ff))
		case "bottomColor":
			p.bottomColor = getRGBAColor(parseIntHex(first, 0x0000ffff))
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

func parseIntHex(str string, defaultVal int) int {
	// must use 64 bit insted of 32 bit to avoid parsing error
	i, err := strconv.ParseInt(str, 16, 64)
	if err != nil {
		return defaultVal
	}
	return int(i)
}

func getRGBAColor(v int) color.RGBA {
	return color.RGBA{
		uint8(v & 0xff000000 >> 24),
		uint8(v & 0x00ff0000 >> 16),
		uint8(v & 0x0000ff00 >> 8),
		uint8(v & 0x000000ff),
	}
}

func compare(a, b param) bool {
	if &a == &b {
		return true
	}

	if a.width != b.width {
		return false
	}
	if a.height != b.height {
		return false
	}
	if a.topColor != b.topColor {
		return false
	}
	if a.bottomColor != b.bottomColor {
		return false
	}
	return true
}
