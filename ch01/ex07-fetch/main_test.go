package main

import (
	"bytes"
	"testing"
)

func TestFetch(t *testing.T) {
	out = new(bytes.Buffer) // capture output
	url := "http://gopl.io"
	if err := fetch(url); err != nil {
		t.Errorf("%s failed:%s", err)
	}
	got := out.(*bytes.Buffer).String()
	if len(got) == 0 {
		t.Error("empty")
	}
}

func TestFetchFailed(t *testing.T) {
	var tests = []struct {
		url string
		err string
	}{
		{"http://gopla.io", "fetch: Get http://gopla.io: dial tcp: lookup gopla.io: no such host"},
		{"httpa://gopl.io", "fetch: Get httpa://gopl.io: unsupported protocol scheme \"httpa\""},
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
