package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	in  io.Reader = os.Stdin
	out io.Writer = os.Stdout
)

type elementProperty struct {
	name  string
	id    string
	class string
}

func (ep elementProperty) String() string {
	str := ep.name

	var attr string
	if ep.id != "" {
		attr += fmt.Sprintf("id:%s,", ep.id)
	}
	if ep.class != "" {
		attr += fmt.Sprintf("class:%s,", ep.class)
	}
	attr = strings.TrimSuffix(attr, ",")
	if attr != "" {
		attr = "[" + attr + "]"
	}

	return str + attr
}

func main() {
	selectXml(os.Args[1:])
}

func selectXml(args []string) {
	dec := xml.NewDecoder(in)
	var stack []elementProperty // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			ep := elementProperty{name: tok.Name.Local}
			for _, attr := range tok.Attr {
				if attr.Name.Local == "id" {
					ep.id = attr.Value
				}
				if attr.Name.Local == "class" {
					ep.class = attr.Value
				}
			}
			stack = append(stack, ep) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, args) {
				var str string
				for i, v := range stack {
					if i > 0 {
						str += " "
					}
					str += v.String()
				}
				fmt.Fprintf(out, "%s: %s\n", str, tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []elementProperty, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0].name == y[0] || x[0].id == y[0] || x[0].class == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
