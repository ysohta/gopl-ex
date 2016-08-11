package main

import (
	"testing"
)

func TestRemoveNeiboringDuplication(t *testing.T) {
	tests := []struct {
		s    []string
		want []string
	}{
		{
			[]string{"a", "a", "b", "c", "c", "c"},
			[]string{"a", "b", "c"},
		}, {
			[]string{"a", "a", "a"},
			[]string{"a"},
		}, {
			[]string{},
			[]string{},
		},
	}

	for _, test := range tests {
		got := removeNeiboringDuplication(test.s)
		if !equal(got, test.want) {
			t.Errorf("Expected:%v Actual:%v", test.want, got)
		}
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
