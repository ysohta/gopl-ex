package sexpr

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	type Person struct {
		Name    string `sexpr:"n"`
		Age     int    `sexpr:"a"`
		Address string
	}
	for _, test := range []struct {
		p    Person
		want string
	}{
		{
			Person{"joe", 30, "hoge"},
			`((n "joe") (a 30) (Address "hoge"))`,
		},
		{
			Person{},
			`((n "") (a 0) (Address ""))`,
		},
	} {
		// Encode it
		data, err := Marshal(&test.p)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if string(data) != test.want {
			t.Errorf("want:%q got:%q", test.want, data)
		}
	}
}

func TestUnmarshal(t *testing.T) {
	type Person struct {
		Name    string `sexpr:"n"`
		Age     int    `sexpr:"a"`
		Address string
	}
	for _, test := range []struct {
		data string
		want Person
	}{
		{
			`((n "joe") (a 30) (Address "hoge"))`,
			Person{"joe", 30, "hoge"},
		},
		{
			`((n "") (a 0) (Address ""))`,
			Person{},
		},
	} {
		// Decode it
		var p Person
		if err := Unmarshal([]byte(test.data), &p); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}

		// Check equality.
		if !reflect.DeepEqual(p, test.want) {
			t.Fatal("want:%v got:%v", test.want, p)
		}
	}
}

// Test verifies that encoding and decoding a complex data value
// produces an equal result.
//
// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
//
// 	$ go test -v gopl.io/ch12/sexpr
//
func Test(t *testing.T) {
	type Movie struct {
		Title    string            `sexpr:"t"`
		Subtitle string            `sexpr:"st"`
		Year     int               `sexpr:"y"`
		Actor    map[string]string `sexpr:"a"`
		Oscars   []string          `sexpr:"o"`
		Sequel   *string           `sexpr:"seq"`
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
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

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Decode it
	var movie Movie
	if err := Unmarshal(data, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	// Check equality.
	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	// Pretty-print it:
	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}
