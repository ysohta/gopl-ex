package strings

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"canoe", "ocean", true},
		{"canoe", "oceam", false},
		{"canoe", "cano", false}, // reject partial
		{"cano", "canoe", false},
		{"アナグラム", "グアムナラ", true},          // accept Japanese
		{"Canoe", "Ocean", true},          // case insensitive
		{"dormitory", "dirty room", true}, // ignore space
	}

	for _, test := range tests {
		got := anagram(test.s1, test.s2)
		if test.want != got {
			t.Errorf("'%s' '%s' Actual: %v\tExpected: %v", test.s1, test.s2, got, test.want)
		}
	}
}
