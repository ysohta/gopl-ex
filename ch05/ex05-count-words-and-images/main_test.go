package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	tests := []struct {
		s             string
		words, images int
	}{
		{
			"<a href='link1'>This is link.</a><img src='sample.png'>",
			3, 1,
		},
	}

	for _, test := range tests {
		n, err := html.Parse(strings.NewReader(test.s))
		if err != nil {
			t.Errorf("parse failure: %v", err)
		}

		w, i := countWordsAndImages(n)
		if w != test.words {
			t.Errorf("expected:%d actual:%d", test.words, w)
		}
		if i != test.images {
			t.Errorf("expected:%d actual:%d", test.images, i)
		}
	}
}
