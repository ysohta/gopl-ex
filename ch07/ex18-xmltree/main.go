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

func main() {
	root := parse(os.Args[1:])
	for _, n := range root.Children {
		fmt.Fprintf(out, "%s", n)
	}
}

func parse(args []string) Element {
	dec := xml.NewDecoder(in)
	parent := Element{Children: []Node{}}
	stack := []*Element{&parent}
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmltree: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			el := NewElement()
			el.Type.Local = tok.Name.Local
			el.Attr = tok.Attr

			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, el)

			stack = append(stack, el) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			trimed := strings.TrimSpace(string(tok))
			if trimed == "" {
				continue
			}
			el := CharData(trimed)
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, el)
		}
	}
	return *stack[0]
}
