package main

import "testing"

func TestWordCounterWrite(t *testing.T) {
	tests := []struct {
		data []string
		want int
	}{
		{
			[]string{"hello world"},
			2,
		},
		{
			[]string{"  can   trim  "},
			2,
		},
		{
			[]string{"there\nare\nlines"},
			3,
		},
	}

	for _, test := range tests {
		var c WordCounter

		for _, s := range test.data {
			c.Write([]byte(s))
		}

		got := int(c)
		if got != test.want {
			t.Errorf("expected=%v actual=%v", test.want, got)
		}

	}
}
