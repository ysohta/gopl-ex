package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	connect()
}

func connect() {
	listener, err := net.Listen("tcp", ":21")
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

func connectDataTransfer() (port int, err error) {
	//  open random port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	_, p, err := net.SplitHostPort(l.Addr().String())
	if err != nil {
		return 0, err
	}
	port, err = strconv.Atoi(p)
	if err != nil {
		return 0, err
	}

	go func(l net.Listener) {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			return
		}

		log.Printf("connected port:%d", port)

		go handleDataTransferConn(conn)
	}(l)

	return port, nil
}

func connectDataTransferPort(port int) (err error) {
	addr := fmt.Sprintf(":%d", port)
	l, err := net.Listen("tcp", addr)

	go func(l net.Listener) {
		for {
			log.Printf("connecting port:%d", port)

			conn, err := l.Accept()
			if err != nil {
				log.Print(err)
				return
			}

			go handleDataTransferConn(conn)
		}
	}(l)

	return nil
}

func handleConn(c net.Conn) {
	defer c.Close()

	sdr := sender{c}

	sdr.sendReplyCode(ReplyCodeOkay)

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
		f := GetCommand(ct)
		if len(strs) > 1 {
			f(sdr, strs[1:]...)
		} else {
			f(sdr)
		}
	}
}

func handleDataTransferConn(c net.Conn) {
	defer c.Close()

	sdr := sender{c}

	dataTransfer = make(chan string)

	for {
		select {
		case data, ok := <-dataTransfer:
			if !ok {
				break
			}
			sdr.sendData(data)
		case <-importing:

			p := make([]byte, 1024)
			r := bufio.NewReader(sdr.c)
			n, err := r.Read(p)
			if err != nil {
				if err != io.EOF {
					// unknown error occurs
					log.Printf("read error: %s", err)
				}

				transferred <- "done"
				return
			}

			dataImport <- p[:n]

			log.Print("transferred")
			transferred <- "done"
			return
		}
	}
}
