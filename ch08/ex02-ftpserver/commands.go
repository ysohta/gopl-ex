package main

import (
	"bufio"
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
	RETR CommandType = "RETR"
	MDTM CommandType = "MDTM"
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
		RETR: retr,
		MDTM: mdtm,
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

	msg := fmt.Sprintf("\"%s\"", dir)
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

	sdr.sendReplyCodeWithMessage(ReplyCodeFileActionComplete, "CWD command successful")
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

	// log.Print(port)
	// sdr.sendReplyCodeWithMessage(ReplyCodeCommandUnrecognized, "EPSV command is not implemented")

	msg := fmt.Sprintf("Entering Extended Passive Mode (|||%d|)", port)
	sdr.sendReplyCodeWithMessage(ReplyCodeEnteringEpsv, msg)
}

func port(sdr sender, req ...string) {
	sdr.sendReplyCode(ReplyCodeNotImplemented)
}

func list(sdr sender, req ...string) {
	sdr.sendReplyCodeWithMessage(ReplyCodeFileStatusOkay, "Opening ASCII mode data connection for file list")

	out, err := exec.Command("ls", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range strings.Split(string(out), "\n") {
		if s == "" || strings.HasPrefix(s, "total") {
			continue
		}
		dataTransfer <- fmt.Sprintf("%s\r\n", s)
	}

	close(dataTransfer)

	<-transferred

	sdr.sendReplyCodeWithMessage(ReplyCodeCloseDataConnection, "Transfer complete")
}

func retr(sdr sender, req ...string) {
	if len(req) < 1 {
		sdr.sendReplyCode(ReplyCodeParameterError)
		return
	}

	pathname := req[0]
	fi, err := os.Stat(pathname)
	if err != nil {
		sdr.sendReplyCodeWithMessage(ReplyCodeFileUnavailable, "unable to open")
		return
	}
	msg := fmt.Sprintf("Opening BINARY mode data connection for %s (%d bytes)", fi.Name(), fi.Size())
	sdr.sendReplyCodeWithMessage(ReplyCodeFileStatusOkay, msg)

	f, err := os.Open(pathname)
	if err != nil {
		sdr.sendReplyCodeWithMessage(ReplyCodeFileUnavailable, "unable to open")
		return
	}
	defer f.Close()

	p := make([]byte, 1024)
	r := bufio.NewReader(f)
	for {
		n, err := r.Read(p)
		if err != nil {
			break
		}
		dataTransfer <- string(p[:n])
	}

	// for _, s := range strings.Split(string(out), "\n") {
	// 	if s == "" || strings.HasPrefix(s, "total") {
	// 		continue
	// 	}
	// 	dataTransfer <- fmt.Sprintf("%s\r\n", s)
	// }

	close(dataTransfer)

	<-transferred

	sdr.sendReplyCodeWithMessage(ReplyCodeCloseDataConnection, "Transfer complete")
}

func mdtm(sdr sender, req ...string) {
	if len(req) < 1 {
		sdr.sendReplyCode(ReplyCodeParameterError)
		return
	}
	fi, err := os.Stat(req[0])
	if err != nil {
		sdr.sendReplyCodeWithMessage(ReplyCodeFileUnavailable, "unable to open")
		return
	}

	t := fi.ModTime().Format("20060102150405")
	sdr.sendReplyCodeWithMessage(ReplyCodeFileStatus, t)
}
