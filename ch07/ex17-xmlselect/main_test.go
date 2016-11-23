package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSelectXml(t *testing.T) {
	tests := []struct {
		tags []string
		in   string
		want string
	}{
		{
			[]string{"a", "b", "c"},
			"<a><b><c>inside c</c></b><b><d>inside d</d></b><b></b></a>",
			"a b c: inside c\n",
		},
		{
			[]string{"a", "myid", "myclass"},
			`<a>
			  <b>
			    <c class="myclass">without id</c>
			  </b>
			  <b id="myid">
			    <c class="myclass">inside c</c>
			    <c>without class</c>
			  </b>
			</a>
			`,
			"a b[id:myid] c[class:myclass]: inside c\n",
		},
		{
			[]string{"myclass"},
			`<a>
			  <b>
			    <c class="myclass">without id</c>
			  </b>
			  <b id="myid">
			    <c class="myclass">inside c</c>
			    <c>without class</c>
			  </b>
			</a>
			`,
			"a b c[class:myclass]: without id\na b[id:myid] c[class:myclass]: inside c\n",
		},
	}

	for _, test := range tests {
		buf := bytes.NewBufferString("")
		in = strings.NewReader(test.in)
		out = buf

		selectXml(test.tags)

		got := buf.String()
		if got != test.want {
			t.Errorf("got: %q want: %q", got, test.want)
		}
	}
}
