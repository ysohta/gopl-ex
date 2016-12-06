package main

import (
	"fmt"
	"log"
	"net"
)

const (
	minPort = 0xc000
	maxPort = 0xffff
)

func ListenEphemeralPort(network string) (net.Listener, int, error) {
	//	for p := minPort; p <= maxPort; p++ {
	laddr := fmt.Sprintf(":%d", 0)
	log.Printf("%s", laddr)
	if l, err := net.Listen(network, laddr); err == nil {
		log.Printf("%s", l.Addr().String())

		return l, 0, nil
		//	}
	}
	return nil, 0, fmt.Errorf("cannot find open port")
}
