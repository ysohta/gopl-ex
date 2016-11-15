package wordcounter

import (
	"bufio"
	"strings"
)

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	var cnt int
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	for scanner.Scan() {
		cnt++
	}

	*c += LineCounter(cnt)
	return cnt, scanner.Err()
}
