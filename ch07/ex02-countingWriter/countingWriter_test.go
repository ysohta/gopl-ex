package countingWriter

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	tests := []struct {
		data []string
		cnt  int64
		want string
	}{
		{
			[]string{
				"hello", " ", "world",
			},
			11,
			"hello world",
		},
		{
			[]string{
				"こんにちは", "世界",
			},
			21,
			"こんにちは世界",
		},
		{
			[]string{},
			0,
			"",
		},
	}

	for _, test := range tests {
		w := bytes.NewBufferString("")
		cw, cnt := CountingWriter(w)

		for _, d := range test.data {
			_, err := cw.Write([]byte(d))
			if err != nil {
				t.Errorf("error:%s", err)
			}
		}

		if *cnt != test.cnt {
			t.Errorf("expected:%d actual:%d", test.cnt, *cnt)
		}
		if w.String() != test.want {
			t.Errorf("expected:%s actual:%s", test.want, w)
		}
	}
}
