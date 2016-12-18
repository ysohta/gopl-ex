package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	steps    = flag.Int("steps", 5, "number of steps of piepeline")
	messages = flag.Int("messages", 1, "number of messages to send")
)

func generate(out chan<- int) {
	for x := 0; x < *messages; x++ {
		out <- x
	}
	close(out)
}

func pass(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v
	}
	close(out)
}

func consume(in <-chan int) {
	for _ = range in {
		// do nothing
	}
}

func main() {
	flag.Parse()

	var prev, next chan int
	prev = make(chan int)

	// create piepeline
	t0 := time.Now()
	for i := 0; i < *steps; i++ {
		next = make(chan int)
		go pass(next, prev)
		prev = next
	}

	// start to send
	t1 := time.Now()
	go generate(prev)
	consume(next)

	t2 := time.Now()

	fmt.Printf("%v for initialization.\n", t1.Sub(t0))
	fmt.Printf("%v to run.\n", t2.Sub(t1))
}
