package wordcounter

import (
	"bufio"
	"strings"
)

type WordCounter int

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
