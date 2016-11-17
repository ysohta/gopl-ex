package temp

import (
	"fmt"
	"testing"
)

func TestStringImplicit(t *testing.T) {
	var c Celsius
	var got string
	var want string

	c = Celsius(123)

	want = "123°C"
	got = fmt.Sprintf("%s", c) // String() is called
	if got != want {
		t.Errorf("expected:%s actual:%s", want, got)
	}

	want = "123.00"
	got = fmt.Sprintf("%.2f", c) // print as float64
	if got != want {
		t.Errorf("expected:%s actual:%s", want, got)
	}
}
