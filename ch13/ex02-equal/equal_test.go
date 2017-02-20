package equal

import "testing"

type link struct {
	value string
	tail  *link
}

func TestHasCycle(t *testing.T) {
	// Circular linked lists a -> b -> a and c -> c.
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c

	d1, d2 := &link{value: "d"}, &link{value: "d"}
	d1.tail = d2

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		{a, true},
		{b, true},
		{c, true},
		{d1, false},
		{d2, false},
		{"str", false},
		{3, false},
		{[]*link{d1, a}, true},
		{map[string]*link{"d1": d1, "a": a}, true},
	} {
		got := HasCycle(test.x)
		if test.want != got {
			t.Errorf("HasCycle(%v) want:%t got:%t", test.x, test.want, got)
		}
	}
}
