package main

import (
	"bytes"
	"testing"
)

func TestFetch(t *testing.T) {
	out = new(bytes.Buffer) // capture output
	var tests = []struct {
		url string
	}{
		{"http://gopl.io"},
		{"gopl.io"},
		{"https://golang.org/"},
		{"golang.org/"},
	}
	for _, test := range tests {
		if err := fetch(test.url); err != nil {
			t.Errorf("%s failed:%s", err)
		}

		got := out.(*bytes.Buffer).String()
		if len(got) == 0 {
			t.Error("empty")
		}
	}
}

func TestFetchFailed(t *testing.T) {
	var tests = []struct {
		url string
	}{
		{"http://gopla.io"},
		{"ftp://gopla.io"},
	}
	for _, test := range tests {
		err := fetch(test.url)
		if err == nil {
			t.Error("error is missing\n")
		}
	}
}
