package main

import (
	"testing"
)

func TestFtToM(t *testing.T) {
	tests := []struct {
		ft    Feet
		want  Meter
		delta Meter
	}{
		{0, 0, 0.000},
		{1, 0.3048, 0.000},
	}

	for _, test := range tests {
		got := FtToM(test.ft)
		diff := got - test.want
		if diff < -test.delta || diff > test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestMToFt(t *testing.T) {
	tests := []struct {
		m     Meter
		want  Feet
		delta Feet
	}{
		{0, 0, 0.000},
		{0.3048, 1, 0.000},
	}

	for _, test := range tests {
		got := MToFt(test.m)
		diff := got - test.want
		if diff < -test.delta || diff > test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestFeetString(t *testing.T) {
	got := Feet(100).String()
	want := "100ft"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}

func TestMeterString(t *testing.T) {
	got := Meter(100).String()
	want := "100m"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}
