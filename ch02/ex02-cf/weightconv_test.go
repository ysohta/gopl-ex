package main

import (
	"testing"
)

func TestLbToKg(t *testing.T) {
	tests := []struct {
		lb    Pound
		want  Kilogram
		delta Kilogram
	}{
		{0, 0, 0.000},
		{1, 2.20462, 0.001},
	}

	for _, test := range tests {
		got := LbToKg(test.lb)
		diff := got - test.want
		if diff < -test.delta || diff > test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestKgToLb(t *testing.T) {
	tests := []struct {
		kg    Kilogram
		want  Pound
		delta Pound
	}{
		{0, 0, 0.000},
		{1, 0.45359237, 0.000},
	}

	for _, test := range tests {
		got := KgToLb(test.kg)
		diff := got - test.want
		if diff < -test.delta || diff > test.delta {
			t.Errorf("Expected:%g Actual:%g", test.want, got)
		}
	}
}

func TestPoundString(t *testing.T) {
	got := Pound(100).String()
	want := "100lb"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}

func TestKilogramString(t *testing.T) {
	got := Kilogram(100).String()
	want := "100kg"
	if got != want {
		t.Errorf("Expected:%g Actual:%g", want, got)
	}
}
