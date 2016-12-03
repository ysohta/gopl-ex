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

	for {
		s, err := bufio.NewReader(c).ReadString('\n')
		if err == io.EOF {
			log.Print("reached EOF")
			return
		}
		if err != nil {
			log.Print(err)
			return
		}

		s = strings.TrimSpace(s)
		log.Printf("c> %s", s)

		strs := strings.Split(s, Separator)
		cmd := GetCommand(CommandType(strs[0]))

		var res string
		if len(strs) > 1 {
			res = cmd(strs[1:]...)
		} else {
			res = cmd()
		}

		io.WriteString(c, res)
		log.Printf("s> %s", res)
	}
}
