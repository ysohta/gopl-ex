package main

import "testing"

func TestJoin(t *testing.T) {
	want := "foo bar"
	got := join("foo", " ", "bar")
	if got != want {
		t.Errorf("expected:%s actual:%s", want, got)
	}
}
