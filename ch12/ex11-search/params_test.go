package main

import "testing"

type data struct {
	Labels     []string `http:"l"`
	MaxResults int      `http:"max"`
	Exact      bool     `http:"x"`
	NotShown   string
}

func TestPack(t *testing.T) {
	for _, test := range []struct {
		dt   data
		want string
		err  error
	}{
		{
			data{Labels: []string{"golang", "programming"}, MaxResults: 2, Exact: true},
			"l=golang&l=programming&max=2&x=true",
			nil,
		},
		{
			data{NotShown: "str"},
			"max=0&x=false",
			nil,
		},
	} {

		got, _ := Pack(&test.dt)
		if got != test.want {
			t.Errorf("want:%q got:%q", test.want, got)
		}
	}
}

/*
//!+output
$ go build gopl.io/ch12/search
$ ./search &
$ ./fetch 'http://localhost:12345/search'
Search: {Labels:[] MaxResults:10 Exact:false}
$ ./fetch 'http://localhost:12345/search?l=golang&l=programming'
Search: {Labels:[golang programming] MaxResults:10 Exact:false}
$ ./fetch 'http://localhost:12345/search?l=golang&l=programming&max=100'
Search: {Labels:[golang programming] MaxResults:100 Exact:false}
$ ./fetch 'http://localhost:12345/search?x=true&l=golang&l=programming'
Search: {Labels:[golang programming] MaxResults:10 Exact:true}
$ ./fetch 'http://localhost:12345/search?q=hello&x=123'
x: strconv.ParseBool: parsing "123": invalid syntax
$ ./fetch 'http://localhost:12345/search?q=hello&max=lots'
max: strconv.ParseInt: parsing "lots": invalid syntax
//!-output
*/
