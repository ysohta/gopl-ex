package display

import (
	"io"
	"net"
	"os"
	"reflect"
	"sync"
	"testing"
)

func Example_slice() {
	Display("slice", []*int{new(int), nil})
	// Output:
	// Display slice ([]*int):
	// (*slice[0]) = 0
	// slice[1] = nil
}

func Example_nilInterface() {
	var w io.Writer
	Display("w", w)
	// Output:
	// Display w (<nil>):
	// w = invalid
}

func Example_ptrToInterface() {
	var w io.Writer
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w) = nil
}

func Example_struct() {
	Display("x", struct{ x interface{} }{3})
	// Output:
	// Display x (struct { x interface {} }):
	// x.x.type = int
	// x.x.value = 3
}

func Example_interface() {
	var i interface{} = 3
	Display("i", i)
	// Output:
	// Display i (int):
	// i = 3
}

func Example_ptrToInterface2() {
	var i interface{} = 3
	Display("&i", &i)
	// Output:
	// Display &i (*interface {}):
	// (*&i).type = int
	// (*&i).value = 3
}

func Example_array() {
	Display("x", [1]interface{}{3})
	// Output:
	// Display x ([1]interface {}):
	// x[0].type = int
	// x[0].value = 3
}

func Example_map() {
	Display("x", map[string]string{"key": "value"})
	// Output:
	// Display x (map[string]string):
	// x["key"] = "value"
}

func Example_movie() {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("strangelove", strangelove)

	// We don't use an Output: comment since displaying
	// a map is nondeterministic.

	// //!+output
	// Display strangelove (display.Movie):
	// strangelove.Title = "Dr. Strangelove"
	// strangelove.Subtitle = "How I Learned to Stop Worrying and Love the Bomb"
	// strangelove.Year = 1964
	// strangelove.Color = false
	// strangelove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
	// strangelove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
	// strangelove.Actor["Maj. T.J. \"King\" Kong"] = "Slim Pickens"
	// strangelove.Actor["Dr. Strangelove"] = "Peter Sellers"
	// strangelove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
	// strangelove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
	// strangelove.Oscars[0] = "Best Actor (Nomin.)"
	// strangelove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
	// strangelove.Oscars[2] = "Best Director (Nomin.)"
	// strangelove.Oscars[3] = "Best Picture (Nomin.)"
	// strangelove.Sequel = nil
	// //!-output

}

// This test ensures that the program terminates without crashing.
func Test(t *testing.T) {
	// Some other values (YMMV)
	Display("os.Stderr", os.Stderr)
	// Output:
	// Display os.Stderr (*os.File):
	// (*(*os.Stderr).file).fd = 2
	// (*(*os.Stderr).file).name = "/dev/stderr"
	// (*(*os.Stderr).file).nepipe = 0

	var w io.Writer = os.Stderr
	Display("&w", &w)
	// Output:
	// Display &w (*io.Writer):
	// (*&w).type = *os.File
	// (*(*(*&w).value).file).fd = 2
	// (*(*(*&w).value).file).name = "/dev/stderr"
	// (*(*(*&w).value).file).nepipe = 0

	var locker sync.Locker = new(sync.Mutex)
	Display("(&locker)", &locker)
	// Output:
	// Display (&locker) (*sync.Locker):
	// (*(&locker)).type = *sync.Mutex
	// (*(*(&locker)).value).state = 0
	// (*(*(&locker)).value).sema = 0

	Display("locker", locker)
	// Output:
	// Display locker (*sync.Mutex):
	// (*locker).state = 0
	// (*locker).sema = 0
	// (*(&locker)) = nil

	locker = nil
	Display("(&locker)", &locker)
	// Output:
	// Display (&locker) (*sync.Locker):
	// (*(&locker)) = nil

	ips, _ := net.LookupHost("golang.org")
	Display("ips", ips)
	// Output:
	// Display ips ([]string):
	// ips[0] = "173.194.68.141"
	// ips[1] = "2607:f8b0:400d:c06::8d"

	// Even metarecursion!  (YMMV)
	Display("rV", reflect.ValueOf(os.Stderr))
	// Output:
	// Display rV (reflect.Value):
	// (*rV.typ).size = 8
	// (*rV.typ).ptrdata = 8
	// (*rV.typ).hash = 871609668
	// (*rV.typ)._ = 0
	// ...

	// a pointer that points to itself
	type P *P
	var p P
	p = &p
	if false {
		Display("p", p)
		// Output:
		// Display p (display.P):
		// ...stuck, no output...
	}

	// a map that contains itself
	type M map[string]M
	m := make(M)
	m[""] = m
	if false {
		Display("m", m)
		// Output:
		// Display m (display.M):
		// ...stuck, no output...
	}

	// a slice that contains itself
	type S []S
	s := make(S, 1)
	s[0] = s
	if false {
		Display("s", s)
		// Output:
		// Display s (display.S):
		// ...stuck, no output...
	}

	// a linked list that eats its own tail
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}

	Display("c", c)
	// Output:
	// Display c (display.Cycle):
	// c.Value = 42
	// (*c.Tail).Value = 42
	// (*(*c.Tail).Tail).Value = 42
	// (*(*c.Tail).Tail).Tail = *display.Cycle 0xc4200773f0
}
