package main

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const tempFile = "tempFile.txt"

func execute(name, s string) (string, error) {
	var err error
	var b []byte

	// ignore removing result
	err = os.Remove(tempFile)

	data := []byte(s)
	ioutil.WriteFile(tempFile, data, os.ModePerm)

	cmd := exec.Command(name, tempFile)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	b, err = ioutil.ReadFile(tempFile)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
