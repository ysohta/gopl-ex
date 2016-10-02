package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

var f = func(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "%s ", n.Data)
	}
}

func TestForEachNode(t *testing.T) {
	tests := []struct {
		s    string
		pre  func(n *html.Node)
		post func(n *html.Node)
		want string
	}{
		{
			"<html><head></head><body></body></html>",
			f,
			f,
			"html head head body body html ",
		}, {
			"<html><head></head><body><a href='...'>link</a></body></html>",
			startElement,
			endElement,
			"<html>\n" +
				"  <head/>\n" +
				"  <body>\n" +
				"    <a href='...'>\n" +
				"      link\n" +
				"    </a>\n" +
				"  </body>\n" +
				"</html>\n",
		},
	}

	for _, test := range tests {
		out = new(bytes.Buffer) // capture output

		n, err := html.Parse(strings.NewReader(test.s))
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}

		forEachNode(n, test.pre, test.post)

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("Expected:[%v] Actual:[%v]", test.want, got)
		}
	}
}

func TestStartElement(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{
			"<html><head/><body><!-- comment --></body></html>",
			"<html>\n" +
				"  <head/>\n" +
				"    <body>\n" +
				"      <!-- comment -->\n",
		},
		{
			"<p>text</p>",
			"<html>\n" +
				"  <head/>\n" +
				"    <body>\n" +
				"      <p>\n" +
				"        text\n",
		}, {
			"<img></img>",
			"<html>\n" +
				"  <head/>\n" +
				"    <body>\n" +
				"      <img/>\n",
		}, {
			"<a href='...'>link</a>",
			"<html>\n" +
				"  <head/>\n" +
				"    <body>\n" +
				"      <a href='...'>\n" +
				"        link\n",
		},
	}

	for _, test := range tests {
		depth = 0
		out = new(bytes.Buffer) // capture output

		n, err := html.Parse(strings.NewReader(test.s))
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}

		forEachNode(n, startElement, nil)

		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("Expected:[%v] Actual:[%v]", test.want, got)
		}
	}
}
