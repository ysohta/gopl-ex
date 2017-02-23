package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

type Decoder struct {
	lex   *lexer
	first bool
}

func NewDecoder(r io.Reader) *Decoder {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)

	return &Decoder{lex, true}
}

func (dec *Decoder) Decode(out interface{}) (err error) {
	if dec.lex.scan.Peek() == scanner.EOF {
		return io.EOF
	}
	if dec.first {
		dec.lex.next()
		dec.first = false
	}
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()
	read(dec.lex, reflect.ValueOf(out).Elem())
	return nil
}
