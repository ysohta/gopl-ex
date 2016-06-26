// ex03-benchmark-echo prints benchmark test of echo programs.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	trials := 100

	// benchmarking
	timeEcho1 := benchmark(echo1, trials)
	timeEcho3 := benchmark(echo3, trials)

	// print results
	fmt.Printf("echo1\t%d ns/op\n", timeEcho1)
	// => "echo1	8761 ns/op"

	fmt.Printf("echo3\t%d ns/op\n", timeEcho3)
	// => "echo3	2280 ns/op"
}

func benchmark(fEcho func(), trials int) int64 {
	start := time.Now()
	for i := 0; i < trials; i++ {
		fEcho()
	}
	elapsed := time.Since(start).Nanoseconds()
	return elapsed / int64(trials)
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
