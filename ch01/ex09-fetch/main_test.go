package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestFetch(t *testing.T) {
	out = new(bytes.Buffer) // capture output
	var tests = []struct {
		url    string
		status string
	}{
		{"http://gopl.io", "[200 OK]\n"},
		{"gopl.io", "[200 OK]\n"},
		{"https://golang.org/", "[200 OK]\n"},
		{"golang.org/", "[200 OK]\n"},
	}
	for _, test := range tests {
		if err := fetch(test.url); err != nil {
			t.Errorf("%s failed:%s", err)
		}

		got := out.(*bytes.Buffer).String()
		if len(got) == 0 {
			t.Error("empty")
		}

		if !strings.HasPrefix(got, test.status) {
			t.Errorf("does not have prefix:%s\n", test.status)
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
