package main

import "testing"

func TestSet(t *testing.T) {
	tests := []struct {
		s    string
		want Celsius
	}{
		{
			"3C",
			Celsius(3),
		},
		{
			"3°C",
			Celsius(3),
		},
		{
			"41F",
			Celsius(5),
		},
		{
			"41°F",
			Celsius(5),
		},
		{
			"0K",
			Celsius(-273.15),
		},
		{
			"INVALID",
			Celsius(0),
		},
	}

	for _, test := range tests {
		var f celsiusFlag
		f.Set(test.s)
		got := f.Celsius
		if got != test.want {
			t.Errorf("expected:%f actual:%f", test.want, got)
		}
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		c    float64
		want string
	}{
		{
			123,
			"123°C",
		},
		{
			0.0,
			"0°C",
		},
		{
			-273.15,
			"-273.15°C",
		},
	}

	for _, test := range tests {
		c := Celsius(test.c)
		got := c.String()
		if got != test.want {
			t.Errorf("expected:%s actual:%s", test.want, got)
		}
	}
}
