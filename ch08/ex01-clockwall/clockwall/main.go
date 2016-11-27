package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var (
	area map[string]string = map[string]string{}
)

func init() {
	for _, arg := range os.Args[1:] {
		strs := strings.SplitN(arg, "=", 2)
		area[strs[0]] = strs[1]
	}
}

func main() {
	num := len(area)
	ch := make(chan string, num)

	cities := []string{}
	for city, host := range area {
		cities = append(cities, city)
		fetchTime(ch, city, host)
	}

	var cnt int
	times := map[string]string{}
	for {
		time := <-ch

		strs := strings.SplitN(time, ":", 2)
		times[strs[0]] = strs[1]

		if cnt%num == 0 {
			clock := "|"
			for _, c := range cities {
				clock += fmt.Sprintf("%s %s|", c, times[c])
			}
			fmt.Printf("\r%s", clock)

		}
	}
}

func fetchTime(ch chan string, city, host string) {
	go func() {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			ch <- fmt.Sprintf("%s:%s", city, scanner.Text())
		}
	}()
}
