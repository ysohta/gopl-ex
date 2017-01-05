package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	reset := make(chan struct{})

	go func() {
		tick := time.Tick(1 * time.Second)
		countdown := 10
		for countdown > 0 {
			fmt.Println(countdown)
			select {
			case <-tick:
				countdown--
			case <-reset:
				countdown = 10
			}
		}
		c.Close()
	}()

	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1*time.Second)

		// reset countdown
		reset <- struct{}{}
	}

	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn)
	}
}
