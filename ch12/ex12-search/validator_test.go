package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateByName(t *testing.T) {
	for _, test := range []struct {
		name  string
		param interface{}
		want  bool
		err   error
	}{
		{"CheckEmail", "j@c", true, nil},
		{"CheckEmail", "j@c@s", false, nil},
		{"CheckEmail", "withoutAt", false, nil},
		{"CheckCreditCard", "1234123412341234", true, nil},
		{"CheckCreditCard", "12341234123412345", false, nil},
		{"CheckCreditCard", "123412341234123", false, nil},
		{"hoge", "fuga", false, errors.New("invalid name:hoge")},
	} {
		got, err := validateByName(test.name, test.param)
		if test.want != got {
			t.Errorf("param:%q want:%t got:%t", test.param, test.want, got)
		}

		if fmt.Sprint(test.err) != fmt.Sprint(err) {
			t.Errorf("want:%q got:%q", test.err, err)
		}
	}
}
