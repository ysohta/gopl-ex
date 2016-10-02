package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var searching string
var depth int
var out io.Writer = os.Stdout

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "missing arguments")
		os.Exit(1)
	}

	url := os.Args[1]
	id := os.Args[2]

	err := findElementByID(url, id)
	if err != nil {
		err = fmt.Errorf("findElement: %s", err)
		os.Exit(1)
	}
}

func findElementByID(url, id string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return err
	}
	if n := ElementByID(doc, id); n != nil {
		fmt.Fprintf(out, "found id=%s\n", id)
	} else {
		fmt.Fprintf(out, "not found id=%s\n", id)
	}
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	searching = id
	return forEachNode(doc, startElement, endElement)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if found := forEachNode(c, pre, post); found != nil {
			return found
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}
	return nil
}

func startElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		s := n.Data
		var found bool
		for _, a := range n.Attr {
			s += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
			if a.Key == "id" && a.Val == searching {
				found = true
			}
		}
		fmt.Fprintf(out, "%*s<%s>\n", depth*2, "", s)
		if found {
			return false
		}

		depth++
	}
	return true
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
	return true
}
