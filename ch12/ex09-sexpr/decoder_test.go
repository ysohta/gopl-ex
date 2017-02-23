package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	for _, test := range []struct {
		in string
	}{
		{`(11 2 3)`},
		{`(key "value")`},
		{`(arr ("1" "2"))`},
	} {
		buf := bytes.NewBufferString(test.in)
		dec := NewDecoder(buf)
		var buffer bytes.Buffer
		var last reflect.Type
		for {
			tok, err := dec.Token()
			if err == io.EOF {
				break
			} else if err != nil {
				t.Fatalf("error found:%v", err)
			}
			switch tok := tok.(type) {
			case Symbol:
				buffer.WriteString(fmt.Sprintf("%s ", tok.Name))
			case String:
				if last == reflect.TypeOf(tok) {
					buffer.WriteString(" ")
				}
				buffer.WriteString(fmt.Sprintf("%q", tok.Value))
			case Int:
				if last == reflect.TypeOf(tok) {
					buffer.WriteString(" ")
				}
				buffer.WriteString(fmt.Sprintf("%d", tok.Value))
			case StartList:
				if last == reflect.TypeOf(EndList{}) {
					buffer.WriteString(" ")
				}
				buffer.WriteString("(")
			case EndList:
				buffer.WriteString(")")
			}
			last = reflect.TypeOf(tok)
		}

		got := buffer.String()
		if got != test.in {
			t.Errorf("want:%q got:%q", test.in, got)
		}
	}
}
