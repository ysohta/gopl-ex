package main

import (
	"testing"
)

func TestRepaceSpace(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"a b\tc",
			"a b c",
		}, {
			"a\r\nb",
			"a b",
		}, {
			"\\n\n\\n",
			"\\n \\n",
		}, {
			"全角　スペース",
			"全角 スペース",
		}, {
			"",
			"",
		},
	}

	for _, test := range tests {
		b := []byte(test.s)
		b = replaceSpace(b)
		got := string(b)
		if got != test.want {
			t.Errorf("Expected:%v Actual:%v", test.want, got)
		}
	}
}
