package split

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a::b::c", "::", []string{"a", "b", "c"}},
		{":a:b:c:", ":", []string{"", "a", "b", "c", ""}},
		{"a:b:c", "", []string{"a", ":", "b", ":", "c"}},
	}

	for _, test := range tests {
		got := strings.Split(test.s, test.sep)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got:%v want:%v", got, test.want)
		}
	}
}
