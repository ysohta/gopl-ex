package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

type sender struct {
	c net.Conn
}

func (s *sender) sendReplyCode(rep ReplyCode) (err error) {
	res := fmt.Sprintf("%d\n", rep)
	_, err = io.WriteString(s.c, res)
	log.Printf("s> %s", res)
	return err
}

func (s *sender) sendReplyCodeWithMessage(rep ReplyCode, msg string) (err error) {
	res := fmt.Sprintf("%d %s\n", rep, msg)
	_, err = io.WriteString(s.c, res)
	log.Printf("s> %s", res)
	return err
}

func (s *sender) sendMessage(msg string) (err error) {
	res := fmt.Sprintf("%s\n", msg)
	_, err = io.WriteString(s.c, res)
	log.Printf("s> %s", res)
	return err
}

func (s *sender) sendData(data string) (err error) {
	_, err = io.WriteString(s.c, data)
	log.Printf("s> %s", data)
	return err
}
