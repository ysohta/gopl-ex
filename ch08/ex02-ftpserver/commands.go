package main

import (
	"fmt"
	"log"
	"os"
)

const (
	Separator = " "
)

type CommandType string

const (
	USER CommandType = "USER"
	PASS CommandType = "PASS"
	SYST CommandType = "SYST"
	FEAT CommandType = "FEAT"
	QUIT CommandType = "QUIT"
	PWD  CommandType = "PWD"
	CWD  CommandType = "CWD"
)

var (
	Commands map[CommandType]Command = map[CommandType]Command{
		USER: user,
		PASS: pass,
		FEAT: feat,
		QUIT: quit,
		PWD:  pwd,
		CWD:  cwd,
	}
)

func Validate(cmd CommandType) bool {
	_, ok := Commands[cmd]
	return ok
}

func GetCommand(cmd CommandType) Command {
	log.Print(fmt.Sprintf("%q", cmd))
	if Validate(cmd) {
		return Commands[cmd]
	}

	// TODO: returns 500
	return okay
}

type Command func(req ...string) string

func okay(req ...string) string {
	return fmt.Sprintf("%d\n", ReplyCodeOkay)
}

func user(req ...string) string {
	return fmt.Sprintf("%d\n", ReplyCodeNeedAccount)
}

func pass(req ...string) string {
	return fmt.Sprintf("%d Welcome to FTP server.\n", ReplyCodeUserLoggedIn)
}

func feat(req ...string) string {
	return fmt.Sprintf("%d\n", ReplyCodeNotImplemented)
}

func quit(req ...string) string {
	return fmt.Sprintf("%d\n", ReplyCodeCloeseConnection)
}

func pwd(req ...string) string {
	var dir string
	var err error
	if dir, err = os.Getwd(); err != nil {
		return fmt.Sprintf("%d unaccessible.\n", ReplyCodeFileUnavailable)
	}

	return fmt.Sprintf("%d \"%s\"\n", ReplyCodePathNameCreated, dir)
}

func cwd(req ...string) string {
	if len(req) < 1 {
		return fmt.Sprintf("%d\n", ReplyCodeParameterError)
	}

	if err := os.Chdir(req[0]); err != nil {
		return fmt.Sprintf("%d %s: No such file or unaccessible.\n", ReplyCodeFileUnavailable, req[0])
	}

	return fmt.Sprintf("%d\n", ReplyCodeFileActionComplete)
}
