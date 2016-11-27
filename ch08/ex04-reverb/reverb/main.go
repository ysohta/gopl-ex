package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	host := "localhost"
	laddr := fmt.Sprintf("%s:%d", host, 8000)

	listener, err := net.Listen("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		var wg sync.WaitGroup
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go func(c net.Conn) {
			input := bufio.NewScanner(c)
			for input.Scan() {
				wg.Add(1)
				go func(c net.Conn, shout string, delay time.Duration) {
					fmt.Fprintln(c, "\t", strings.ToUpper(shout))
					time.Sleep(delay)
					fmt.Fprintln(c, "\t", shout)
					time.Sleep(delay)
					fmt.Fprintln(c, "\t", strings.ToLower(shout))
					wg.Done()
				}(c, input.Text(), 1*time.Second)
			}
			wg.Wait()
			c.Close()
		}(conn)
	}
}
