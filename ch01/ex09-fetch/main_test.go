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
		err string
	}{
		{
			"http://gopla.io",
			"fetch: Get http://gopla.io: dial tcp: lookup gopla.io: no such host"
		},
		{
			"ftp://gopla.io",
			"fetch: Get http://ftp://gopla.io: dial tcp: lookup tcp/: nodename nor servname provided, or not known"
		},
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
