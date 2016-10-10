package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var searching []string
var out io.Writer = os.Stdout

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing arguments")
		os.Exit(1)
	}

	url := os.Args[1]
	names := os.Args[2:]

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Errorf("parsing HTML: %s", err)
		return
	}
	resp.Body.Close()

	for _, n := range ElementByTagName(doc, names...) {
		s := n.Data
		for _, a := range n.Attr {
			s += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
		}
		fmt.Fprintf(out, "<%s/>\n", s)
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	searching = name
	return forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) []*html.Node {
	var nodes []*html.Node
	if pre != nil {
		if pre(n) {
			nodes = append(nodes, n)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := forEachNode(c, pre, post); found != nil {
			nodes = append(nodes, found...)
		}
	}

	if post != nil {
		if post(n) {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		for _, name := range searching {
			if name == n.Data {
				return true
			}
		}
	}
	return false
}

func endElement(n *html.Node) bool {
	return false
}
