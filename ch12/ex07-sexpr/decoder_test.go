package sexpr

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type data struct {
		R int
		S []string
	}

	for _, test := range []struct {
		stream string
		want   []data
	}{
		{
			`((R 2) (S ("A" "B")))
			((R 4) (S ("C")))
			((R 1) (S ("k")))`,
			[]data{
				data{R: 2, S: []string{"A", "B"}},
				data{R: 4, S: []string{"C"}},
				data{R: 1, S: []string{"k"}},
			},
		},
	} {
		r := bytes.NewBufferString(test.stream)
		dec := NewDecoder(r)
		var got []data
		for {
			var dt data
			err := dec.Decode(&dt)
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("unexpected:%v", err)
			}
			got = append(got, dt)
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("want:%v got:%v", test.want, got)
		}
	}
}
