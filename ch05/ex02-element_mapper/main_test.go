package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestMapElements(t *testing.T) {
	tests := []struct {
		s    string
		want map[string]int
	}{
		{
			"<html><head/><body><div><a href='link1'><a href='link2'></div><a href='link3'></body></html>",
			map[string]int{"html": 1, "head": 1, "body": 1, "div": 1, "a": 3},
		},
	}

	for _, test := range tests {
		n, err := html.Parse(strings.NewReader(test.s))
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
		got := mapElements(map[string]int{}, n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Expected:%v Actual:%v", test.want, got)
		}
	}
}
