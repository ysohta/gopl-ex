package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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
	PASV CommandType = "PASV"
	EPSV CommandType = "EPSV"
	LIST CommandType = "LIST"
	PORT CommandType = "PORT"
)

var (
	Commands     map[CommandType]Command
	dataTransfer chan string = make(chan string)
	transferred  chan string = make(chan string)
)

func init() {
	Commands = map[CommandType]Command{
		USER: user,
		PASS: pass,
		FEAT: feat,
		QUIT: quit,
		PWD:  pwd,
		CWD:  cwd,
		PASV: pasv,
		EPSV: epsv,
		LIST: list,
		PORT: port,
	}
}

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

type Command func(sdr sender, req ...string)

func okay(sdr sender, req ...string) {
	sdr.sendReplyCode(ReplyCodeOkay)
}

func user(sdr sender, req ...string) {
	sdr.sendReplyCode(ReplyCodeNeedAccount)
}

func pass(sdr sender, req ...string) {
	sdr.sendReplyCodeWithMessage(ReplyCodeUserLoggedIn, "Welcome to FTP server.")
}

func feat(sdr sender, req ...string) {
	sdr.sendReplyCode(ReplyCodeNotImplemented)
}

func quit(sdr sender, req ...string) {
	sdr.sendReplyCode(ReplyCodeCloeseConnection)
}

func pwd(sdr sender, req ...string) {
	var dir string
	var err error
	if dir, err = os.Getwd(); err != nil {
		sdr.sendReplyCodeWithMessage(ReplyCodeFileUnavailable, "unaccessible")
		return
	}

	msg := fmt.Sprintf("\"%s\"\n", dir)
	sdr.sendReplyCodeWithMessage(ReplyCodePathNameCreated, msg)
}

func cwd(sdr sender, req ...string) {
	if len(req) < 1 {
		sdr.sendReplyCode(ReplyCodeParameterError)
		return
	}

	if err := os.Chdir(req[0]); err != nil {
		msg := fmt.Sprintf("%s: No such file or unaccessible.", req[0])
		sdr.sendReplyCodeWithMessage(ReplyCodeFileUnavailable, msg)
		return
	}

	sdr.sendReplyCode(ReplyCodeFileActionComplete)
}

func pasv(sdr sender, req ...string) {
	port, err := connectDataTransfer()
	if err != nil {
		sdr.sendReplyCode(ReplyCodeParameterError)
		return
	}

	msg := fmt.Sprintf("Entering Passive Mode (127,0,0,1,%d,%d)", port/256, port%256)
	sdr.sendReplyCodeWithMessage(ReplyCodeEnteringPasv, msg)
}

func epsv(sdr sender, req ...string) {
	port, err := connectDataTransfer()
	if err != nil {
		sdr.sendReplyCode(ReplyCodeParameterError)
		return
	}

	log.Print(port)
	sdr.sendReplyCodeWithMessage(ReplyCodeNotImplemented, "EPSV command is not implemented")

	// msg := fmt.Sprintf("Entering Extended Passive Mode (|||%d|)", port)
	// sdr.sendReplyCodeWithMessage(ReplyCodeEnteringEpsv, msg)
}

func port(sdr sender, req ...string) {
	sdr.sendReplyCode(ReplyCodeNotImplemented)
}

func list(sdr sender, req ...string) {
	sdr.sendReplyCodeWithMessage(ReplyCodeFileStatusOkay, "Opening ASCII mode data connection for file list")

	// sample list output
	// strs := []string{
	// 	"lrwxrwxrwx   1 root     root           16 Oct 21  2015 debian -> pub/Linux/debian",
	// 	"lrwxrwxrwx   1 root     root           26 Oct 21  2015 debian-backports -> pub/Linux/debian-backports",
	// 	"lrwxrwxrwx   1 root     root           19 Oct 21  2015 debian-cd -> pub/Linux/debian-cd",
	// 	"lrwxrwxrwx   1 root     root           25 Oct 21  2015 debian-volatile -> pub/Linux/debian-volatile",
	// 	"lrwxrwxrwx   1 root     root           16 Oct 21  2015 gentoo -> pub/Linux/gentoo",
	// 	"drwxr-xr-x  25 ftp-mirror ftp-adm      4096 Oct 18 07:49 pub",
	// }

	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}

	//og.Printf("%q", out)

	// s := strings.Replace(out, "\n", "\r\n", -1)

	for _, s := range strings.Split(string(out), "\n") {
		dataTransfer <- fmt.Sprintf("%s\r\n", s)
	}

	transferred <- "done"

	sdr.sendReplyCode(ReplyCodeCloseDataConnection)
}
