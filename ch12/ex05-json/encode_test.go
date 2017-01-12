package json

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
		err  error
	}{
		{
			[]int{1, 2, 3},
			"[1, 2, 3]",
			nil,
		},
		{
			struct {
				Name string
				List []int
			}{
				"foo", []int{1, 2, 3},
			},
			"{\"Name\": \"foo\", \"List\": [1, 2, 3]}",
			nil,
		},
		{
			map[string]uint{"foo": 55},
			"{\"foo\": 55}",
			nil,
		},
		{
			[2]bool{true, false},
			"[true, false]",
			nil,
		},
		{
			nil,
			"null",
			nil,
		},
	}

	for _, test := range tests {
		got, err := Marshal(test.v)

		if string(got) != test.want {
			t.Errorf("got:%s want:%s", got, test.want)
		}

		if fmt.Sprint(err) != fmt.Sprint(test.err) {
			t.Errorf("got:%s want:%s", err, test.err)
		}
	}
}

func TestMarshalUnmarshal(t *testing.T) {
	tests := []struct {
		v    interface{}
		want interface{}
		got  interface{}
		err1 error
		err2 error
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{},
			nil,
			nil,
		},
		{
			map[string]bool{"T": true},
			map[string]bool{"T": true},
			map[string]bool{},
			nil,
			nil,
		},
		{
			struct{ id int }{3},
			map[string]int{"id": 3}, // struct -> map
			map[string]int{},
			nil,
			nil,
		},
	}

	for _, test := range tests {
		data, err1 := Marshal(test.v)
		if fmt.Sprint(err1) != fmt.Sprint(test.err1) {
			t.Errorf("err1 got:%s want:%s", err1, test.err1)
		}

		err2 := Unmarshal(data, &test.got)
		if fmt.Sprint(test.got) != fmt.Sprint(test.want) {
			t.Errorf("v got:%v want:%v", test.got, test.v)
			t.Errorf("type got:%T want:%T", test.got, test.v)
		}

		if fmt.Sprint(err2) != fmt.Sprint(test.err2) {
			t.Errorf("err2 got:%s want:%s", err2, test.err2)
		}
	}
}
