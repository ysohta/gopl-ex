package main

import (
	"testing"
)

func TestFormatWithIndex(t *testing.T) {
	commands := []string{"cmd", "p1", "p2", "p3"}

	actual := FormatWithIndex(commands)

	expected := "[0]cmd\n[1]p1\n[2]p2\n[3]p3\n"
	if actual != expected {
		t.Error("Actual:", actual, " Expected:", expected)
	}
}
