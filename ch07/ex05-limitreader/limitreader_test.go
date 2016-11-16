package limitreader

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		s     string
		n     int
		size  int
		wantN []int
		wantS []string
	}{
		{
			"hey",
			2,
			10,
			[]int{2},
			[]string{"he"},
		},
		{
			"hey",
			3,
			10,
			[]int{3},
			[]string{"hey"},
		},
		{
			"hey",
			4,
			10,
			[]int{3},
			[]string{"hey"},
		},
		{
			"hey",
			1,
			1,
			[]int{1},
			[]string{"h"},
		},
		{
			"hey",
			3,
			1,
			[]int{1, 1, 1},
			[]string{"h", "e", "y"},
		},
		{
			"hey",
			4,
			1,
			[]int{1, 1, 1},
			[]string{"h", "e", "y"},
		},
	}

	for _, test := range tests {
		r := strings.NewReader(test.s)
		target := LimitReader(r, int64(test.n))

		p := make([]byte, test.size)

		var listN []int
		var listS []string
		for {
			n, err := target.Read(p)
			if err == io.EOF {
				break
			}

			listN = append(listN, n)
			listS = append(listS, string(p[:n]))
		}

		if !reflect.DeepEqual(listN, test.wantN) {
			t.Errorf("expected:%v actual:%v", test.wantN, listN)
		}
		if !reflect.DeepEqual(listS, test.wantS) {
			t.Errorf("expected:%v actual:%v", test.wantS, listS)
		}
	}
}
