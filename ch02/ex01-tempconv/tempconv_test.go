package tempconv

import (
	"fmt"
	"testing"
)

func TestCelciusString(t *testing.T) {
	got := BoilingC.String()
	want := "100°C"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}

func TestFahrenheitString(t *testing.T) {
	got := fmt.Sprint(Fahrenheit(32))
	want := "32°F"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}

func TestKelvinString(t *testing.T) {
	got := fmt.Sprint(Kelvin(0))
	want := "0K"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}

func TestCtoF(t *testing.T) {
	tests := []struct {
		c     Celsius
		want  Fahrenheit
		delta Fahrenheit
	}{
		{AbsoluteZeroC, -459.67, 0.001},
		{FreezingC, 32, 0.001},
		{BoilingC, 212, 0.001},
	}

	for _, test := range tests {
		got := CToF(test.c)
		diff := got - test.want
		if diff <= -test.delta || diff >= test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestFtoC(t *testing.T) {
	tests := []struct {
		f     Fahrenheit
		want  Celsius
		delta Celsius
	}{
		{-459.67, AbsoluteZeroC, 0.001},
		{32, FreezingC, 0.001},
		{212, BoilingC, 0.001},
	}

	for _, test := range tests {
		got := FToC(test.f)
		diff := got - test.want
		if diff <= -test.delta || diff >= test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestKtoF(t *testing.T) {
	tests := []struct {
		k     Kelvin
		want  Fahrenheit
		delta Fahrenheit
	}{
		{0, -459.67, 0.001},
		{273.15, 32, 0.001},
		{373.15, 212, 0.001},
	}

	for _, test := range tests {
		got := KToF(test.k)
		diff := got - test.want
		if diff <= -test.delta || diff >= test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestKtoC(t *testing.T) {
	tests := []struct {
		k     Kelvin
		want  Celsius
		delta Celsius
	}{
		{0, AbsoluteZeroC, 0.001},
		{273.15, FreezingC, 0.001},
		{373.15, BoilingC, 0.001},
	}

	for _, test := range tests {
		got := KToC(test.k)
		diff := got - test.want
		if diff <= -test.delta || diff >= test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestCtoK(t *testing.T) {
	tests := []struct {
		c     Celsius
		want  Kelvin
		delta Kelvin
	}{
		{AbsoluteZeroC, 0, 0.001},
		{FreezingC, 273.15, 0.001},
		{BoilingC, 373.15, 0.001},
	}

	for _, test := range tests {
		got := CToK(test.c)
		diff := got - test.want
		if diff <= -test.delta || diff >= test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestFtoK(t *testing.T) {
	tests := []struct {
		f     Fahrenheit
		want  Kelvin
		delta Kelvin
	}{
		{-459.67, 0, 0.001},
		{32, 273.15, 0.001},
		{212, 373.15, 0.001},
	}

	for _, test := range tests {
		got := FToK(test.f)
		diff := got - test.want
		if diff <= -test.delta || diff >= test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}
