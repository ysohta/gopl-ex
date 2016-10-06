package main

import (
	"reflect"
	"testing"
)

func TestTopoSort(t *testing.T) {
	tests := []struct {
		m       map[string][]string
		want    []string
		acyclic bool
	}{
		{
			map[string][]string{"a": {"b"}, "b": {"a"}},
			nil,
			true,
		}, {
			map[string][]string{"a": {"b"}, "b": {"c"}, "c": {"a"}},
			nil,
			true,
		}, {
			map[string][]string{"a": {"b"}, "b": {"a"}, "c": {"d"}},
			nil,
			true,
		}, {
			map[string][]string{"a": {"c"}, "b": {"c"}},
			[]string{"c", "a", "b"},
			false,
		}, {
			map[string][]string{"a": {"b"}, "b": {"c"}},
			[]string{"c", "b", "a"},
			false,
		}, {
			map[string][]string{"a": {"b", "c"}, "c": {"d", "e"}},
			[]string{"b", "d", "e", "c", "a"},
			false,
		},
	}

	for _, test := range tests {
		got, acyclic := topoSort(test.m)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("expected:%v actual:%v", test.want, got)
		}
		if acyclic != test.acyclic {
			t.Errorf("expected:%v actual:%v", test.acyclic, acyclic)
		}
	}
}
