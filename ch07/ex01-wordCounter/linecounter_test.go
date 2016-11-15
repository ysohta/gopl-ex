package wordcounter

import "testing"

func TestLineCounterWrite(t *testing.T) {
	tests := []struct {
		data []string
		want int
	}{
		{
			[]string{"hello world"},
			1,
		},
		{
			[]string{"  can   trim  "},
			1,
		},
		{
			[]string{"there\nare\nlines"},
			3,
		},
		{
			[]string{""},
			0,
		},
	}

	for _, test := range tests {
		var c LineCounter

		for _, s := range test.data {
			c.Write([]byte(s))
		}

		got := int(c)
		if got != test.want {
			t.Errorf("expected=%v actual=%v", test.want, got)
		}
	}
}
