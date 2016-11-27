package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkGoroutine1(b *testing.B) {
	num = 1
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine2(b *testing.B) {
	num = 2
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine4(b *testing.B) {
	num = 4
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine8(b *testing.B) {
	num = 8
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine16(b *testing.B) {
	num = 16
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine32(b *testing.B) {
	num = 32
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine64(b *testing.B) {
	num = 64
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkGoroutine1024(b *testing.B) {
	num = 1024
	out = ioutil.Discard

	for i := 0; i < b.N; i++ {
		main()
	}
}

func TestMain(t *testing.T) {
	out = ioutil.Discard
	main()
}
