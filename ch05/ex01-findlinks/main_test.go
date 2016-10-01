package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			"<a href='link1'><a href='link2'>",
			[]string{"link1", "link2"},
		}, {
			"<div><a href='link1'><a href='link2'></div><a href='link3'>",
			[]string{"link1", "link2", "link3"},
		},
	}

	for _, test := range tests {
		n, err := html.Parse(strings.NewReader(test.s))
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
		got := visit([]string{}, n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Expected:%v Actual:%v", test.want, got)
		}
	}
}
