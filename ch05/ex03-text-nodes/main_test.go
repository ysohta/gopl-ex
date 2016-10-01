package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestExtractText(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{
			"<html><head/><body><div><a href='link1'>L1</a><a href='link2'>L2</a></div><a href='link3'>L3</a></body></html>",
			[]string{"L1", "L2", "L3"},
		}, {
			"<script>function(){ }</script><a href='link1'>L1</a>",
			[]string{"L1"},
		}, {
			"<style>body {background-color: #fefefe}</style><a href='link1'>L1</a>",
			[]string{"L1"},
		},
	}

	for _, test := range tests {
		n, err := html.Parse(strings.NewReader(test.s))
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}
		got := extractText(nil, n)
		if !reflect.DeepEqual(got, test.want) {
			fmt.Println(len(got), len(test.want))

			t.Errorf("Expected:%v Actual:%v", test.want, got)
		}
	}
}
