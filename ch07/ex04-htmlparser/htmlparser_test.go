package htmlparser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		s     string
		links []string
	}{
		{
			`
			<a href='linkA'>A</a>
			<div>
			  <a href='linkB'>B</a>
			</div>`,
			[]string{"linkA", "linkB"},
		},
	}

	for _, test := range tests {
		n, err := Parse(test.s)
		if err != nil {
			t.Errorf("parse error:%v", err)
		}

		links := visit(nil, n)

		if !reflect.DeepEqual(links, test.links) {
			t.Errorf("expected:%v actual:%v", links, test.links)
		}
	}
}
