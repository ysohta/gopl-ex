package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

type LineCounter int

func main() {
	var c WordCounter
	c.Write([]byte("hello world"))
	fmt.Println(c)

}

// func NewWordCounter(p []byte) *WordCounter {

// 	return WordCounter{len(token), }
// }

func (c *WordCounter) Write(p []byte) (int, error) {
	var cnt int
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		cnt++
	}

	*c += WordCounter(cnt)
	return cnt, scanner.Err()
}

func (c *LineCounter) Write(p []byte) (int, error) {
	var cnt int
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	for scanner.Scan() {
		cnt++
	}

	*c += LineCounter(cnt)
	return cnt, scanner.Err()
}
