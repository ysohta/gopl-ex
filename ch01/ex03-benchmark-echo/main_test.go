package main

import (
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}

func TestEcho1(t *testing.T) {
	actual := benchmark(echo1, 10)

	if actual <= 0 {
		t.Error("Must be positive value")
	}
}

func TestEcho3(t *testing.T) {
	actual := benchmark(echo3, 10)

	if actual <= 0 {
		t.Error("Must be positive value")
	}
}
