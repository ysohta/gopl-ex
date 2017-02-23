package sexpr

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	type data struct {
		R int
		S []string
	}

	for _, test := range []struct {
		s    string
		want data
	}{
		{
			`((R 2) (S ("A" "B")))`,
			data{R: 2, S: []string{"A", "B"}},
		},
	} {
		var dt data
		r := bytes.NewBufferString(test.s)
		dec := NewDecoder(r)
		if err := dec.Decode(&dt); err != nil {
			t.Fatalf("failed Decode:%v", err)
		}

		if !reflect.DeepEqual(dt, test.want) {
			t.Errorf("want:%v got:%v", test.want, dt)
		}
	}
}
