package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCharCount(t *testing.T) {
	tests := []struct {
		in   string
		want result
	}{
		{
			"a,b,c",
			result{
				map[rune]int{
					'a': 1,
					'b': 1,
					'c': 1,
					',': 2,
				},
				[5]int{0, 5, 0, 0, 0},
				0,
			},
		},
		{
			"こんにちは世界",
			result{
				map[rune]int{
					'こ': 1,
					'ん': 1,
					'に': 1,
					'ち': 1,
					'は': 1,
					'世': 1,
					'界': 1,
				},
				[5]int{0, 0, 0, 7, 0},
				0,
			},
		},
		{
			"line\nbreak",
			result{
				map[rune]int{
					'l':  1,
					'i':  1,
					'n':  1,
					'e':  2,
					'b':  1,
					'r':  1,
					'a':  1,
					'k':  1,
					'\n': 1,
				},
				[5]int{0, 10, 0, 0, 0},
				0,
			},
		},
	}

	for _, test := range tests {
		got := charcount(bytes.NewReader([]byte(test.in)))
		if !reflect.DeepEqual(got.counts, test.want.counts) {
			t.Errorf("counts got:%v want:%v", got.counts, test.want.counts)
		}
		if !reflect.DeepEqual(got.utflen, test.want.utflen) {
			t.Errorf("utflen got:%v want:%v", got.utflen, test.want.utflen)
		}
		if got.invalid != test.want.invalid {
			t.Errorf("invalid got:%v want:%v", got.invalid, test.want.invalid)
		}
	}
}

func TestPrintCounts(t *testing.T) {
	tests := []struct {
		res  result
		want string
	}{
		{
			result{
				map[rune]int{
					'a': 1,
				},
				[5]int{0, 1, 0, 0, 0},
				1,
			},
			`rune	count
'a'	1

len	count
1	1
2	0
3	0
4	0

1 invalid UTF-8 characters
`,
		},
	}

	for _, test := range tests {
		var b bytes.Buffer
		out = &b
		printCounts(test.res)
		got := b.String()
		if got != test.want {
			t.Errorf("invalid got:%v want:%v", got, test.want)
		}
	}
}
