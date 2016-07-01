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
		err string
	}{
		{"http://gopla.io", "fetch: Get http://gopla.io: dial tcp: lookup gopla.io: no such host"},
		{"ftp://gopla.io", "fetch: Get http://ftp://gopla.io: dial tcp: lookup tcp/: nodename nor servname provided, or not known"},
	}
	for _, test := range tests {
		err := fetch(test.url)
		actual := err.Error()
		expected := test.err
		if actual != expected {
			t.Errorf("Actual:%s\tExpected:%s", actual, expected)
		}
	}
}
