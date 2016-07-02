package main

import (
	"testing"
)

func TestConvertTemps(t *testing.T) {
	got := convertTemps(32)
	want := "32°F = 0°C, 32°C = 89.6°F"
	if got != want {
		t.Errorf("Actual:%s Expected:%s", got, want)
	}
}

func TestConvertDistances(t *testing.T) {
	got := convertDistances(1000)
	want := "1000ft = 304.8m, 1000m = 3280.839895013123ft"
	if got != want {
		t.Errorf("Actual:%s Expected:%s", got, want)
	}
}

func TestConvertWeights(t *testing.T) {
	got := convertWeights(1000)
	want := "1000lb = 2204.622621848776kg, 1000kg = 453.59237lb"
	if got != want {
		t.Errorf("Actual:%s Expected:%s", got, want)
	}
}
