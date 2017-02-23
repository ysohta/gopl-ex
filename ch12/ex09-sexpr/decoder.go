package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol struct {
	Name string
}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct{}

type EndList struct{}

type Decoder struct {
	scan scanner.Scanner
}

func NewDecoder(r io.Reader) *Decoder {
	dec := Decoder{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	dec.scan.Init(r)
	return &dec
}

func (d *Decoder) Token() (Token, error) {
	for {
		token := d.scan.Scan()
		text := d.scan.TokenText()
		switch token {
		case scanner.EOF:
			return nil, io.EOF
		case scanner.Ident:
			return Symbol{text}, nil
		case scanner.String:
			s, _ := strconv.Unquote(text)
			return String{s}, nil
		case scanner.Int:
			i, _ := strconv.Atoi(text)
			return Int{i}, nil
		case '(':
			return StartList{}, nil
		case ')':
			return EndList{}, nil
		default:
			fmt.Println("default token", token, " text ", d.scan.TokenText())
		}
	}
}
