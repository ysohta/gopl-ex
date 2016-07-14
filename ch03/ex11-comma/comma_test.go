package strings

import (
	"testing"
)

func TestComma(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"12345", "12,345"},
		{"123", "123"},
		{"1234", "1,234"},
		{"123456", "123,456"},
		{"", ""},
		{"123456789", "123,456,789"},
		{"0.1234", "0.1234"},
		{"123.1234", "123.1234"},
		{"1234.1234", "1,234.1234"},
		{"123456.1234", "123,456.1234"},
		{"123456789.1234", "123,456,789.1234"},
	}

	for _, test := range tests {
		got := comma(test.s)
		if test.want != got {
			t.Errorf("Actual: %v\tExpected: %v", got, test.want)
		}
	}
}
