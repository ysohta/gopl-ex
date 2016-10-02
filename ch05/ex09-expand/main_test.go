package main

import (
	"strings"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		s    string
		f    func(string) string
		want string
	}{
		{
			"$foo bar",
			strings.ToUpper,
			"FOO bar",
		}, {
			"Foo $Bar $Baz",
			strings.ToLower,
			"Foo bar baz",
		},
	}

	for _, test := range tests {
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf("Expected:[%v] Actual:[%v]", test.want, got)
		}
	}
}
