package main

import (
	"bytes"
	"fmt"
	"testing"
)

func init() {
	out = new(bytes.Buffer) // capture output
}

func TestMain(t *testing.T) {
	want := fmt.Sprintln("20°C")
	main()
	got := out.(*bytes.Buffer).String()
	if got != want {
		t.Errorf("expected:[%s] actual:[%s]", want, got)
	}
}
