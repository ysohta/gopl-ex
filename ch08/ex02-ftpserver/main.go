package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:21")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()

	io.WriteString(c, okay())

	var ct CommandType

	for ct != QUIT {
		s, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				// unknown error occurs
				log.Printf("read error: %s", err)
			}
			return
		}

		s = strings.TrimSpace(s)
		log.Printf("c> %s", s)

		strs := strings.Split(s, Separator)
		ct = CommandType(strs[0])
		if len(strs) > 1 {
			sendCommand(c, ct, strs[1:]...)
		} else {
			sendCommand(c, ct)
		}
	}
}
