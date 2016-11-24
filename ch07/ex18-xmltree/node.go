package main

import (
	"encoding/xml"
	"fmt"
)

type Node interface {
	String() string
}

type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e Element) String() string {
	str := fmt.Sprintf("<%s", e.Type.Local)

	for _, attr := range e.Attr {
		str += fmt.Sprintf(" %s='%s'", attr.Name.Local, attr.Value)
	}

	str += ">"
	for _, n := range e.Children {
		str += fmt.Sprintf("%s", n)
	}
	str += fmt.Sprintf("</%s>", e.Type.Local)

	return str
}

func NewElement() *Element {
	return &Element{Children: []Node{}}
}
