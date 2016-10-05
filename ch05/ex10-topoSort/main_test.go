package main

import (
	"reflect"
	"testing"
)

func TestTopoSort(t *testing.T) {
	tests := []struct {
		m          map[string][]string
		candidates [][]string
	}{
		{
			map[string][]string{"a": {"b", "c"}},
			[][]string{
				{"b", "c", "a"},
				{"c", "b", "a"},
			},
		}, {
			map[string][]string{"a": {"c"}, "b": {"a"}},
			[][]string{
				{"c", "a", "b"},
			},
		}, {
			map[string][]string{"a": {"c"}, "b": {"c"}, "c": {"d"}},
			[][]string{
				{"d", "c", "a", "b"},
				{"d", "c", "b", "a"},
			},
		},
	}

	for _, test := range tests {
		got := topoSort(test.m)
		var contains bool
		for _, want := range test.candidates {
			if reflect.DeepEqual(got, want) {
				contains = true
			}
		}

		if !contains {
			t.Errorf("%v is not icluded in %v", got, test.candidates)
		}
	}
}
