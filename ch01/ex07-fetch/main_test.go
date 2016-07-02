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
	}{
		{"http://gopla.io"},
		{"httpa://gopl.io"},
	}
	for _, test := range tests {
		err := fetch(test.url)
		if err == nil {
			t.Error("error is missing\n")
		}
	}
}
