package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			"7 - 2 * 3\n",
			"1\n",
		},
		{
			"x + y\n3\n4\n",
			"x: y: 7\n",
		},
		{
			"x % 2\n",
			"parse error: unexpected '%'\n",
		},
		{
			"sqrt(4, 3)\n",
			"check error: call to sqrt has 2 args, want 1\n",
		},
		{
			"x + y\nPi\n",
			"x: parse error: strconv.ParseFloat: parsing \"Pi\": invalid syntax\nx: ",
		},
	}

	for _, test := range tests {
		buf := bytes.NewBufferString("")
		in = strings.NewReader(test.in)
		out = buf

		calc()

		got := buf.String()
		if got != test.want {
			t.Errorf("got:%q, want:%q", got, test.want)
		}

		buf.Reset()
	}
}
