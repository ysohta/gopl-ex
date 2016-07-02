package main

import (
	"net/http"
	"testing"
)

func testParseFormValueToInt(t *testing.T) {
	r := http.Request{
		Form: map[string][]string{
			"cycles": {"20"},
			"res":    {"0.005"},
			"str":    {"str"},
		},
	}
	tests := []struct {
		r    *http.Request
		key  string
		def  int
		want int
	}{
		{&r, "cycles", 5, 20},
		{&r, "size", 100, 100},
		{&r, "res", 100, 100},
		{&r, "str", 100, 100},
	}

	for _, test := range tests {
		got := parseFormValueToInt(test.r, test.key, test.def)
		if got != test.want {
			t.Errorf("Actual: %d\tExpected: %d", got, test.want)
		}
	}
}

func testParseFormValueToFloat64(t *testing.T) {
	r := http.Request{
		Form: map[string][]string{
			"cycles": {"20"},
			"res":    {"0.005"},
			"str":    {"str"},
		},
	}
	tests := []struct {
		r    *http.Request
		key  string
		def  float64
		want float64
	}{
		{&r, "res", 0.001, 0.005},
		{&r, "cycles", 100, 100},
		{&r, "size", 100, 100},
		{&r, "str", 100, 100},
	}

	for _, test := range tests {
		got := parseFormValueToFloat64(test.r, test.key, test.def)
		if got != test.want {
			t.Errorf("Actual: %d\tExpected: %d", got, test.want)
		}
	}
}
