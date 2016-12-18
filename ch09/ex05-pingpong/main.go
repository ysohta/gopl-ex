package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	t = flag.Int("t", 1000, "time[msec] for testing")
)

func main() {
	flag.Parse()

	ping := make(chan uint64)
	pong := make(chan uint64)

	go func() {
		for {
			select {
			case n := <-ping:
				pong <- n + 1
			}
		}
	}()
	go func() {
		for {
			select {
			case n := <-pong:
				ping <- n + 1
			}
		}
	}()

	tick := time.Tick(time.Duration(*t) * time.Millisecond)

	ping <- 0

	for {
		select {
		case <-tick:
			fmt.Println("numbers of rally:", <-ping)
			return
		}
	}
}
