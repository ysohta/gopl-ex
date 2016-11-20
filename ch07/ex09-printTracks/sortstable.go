package main

import "time"

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type comparer interface {
	compare(t1, t2 *Track) int
}
type sortRules struct {
	t     []*Track
	rules []comparer
}

func (x *sortRules) Add(c comparer) {
	//if len(x.rules) == 0 {
	x.rules = append(x.rules, c)
	//	return
	//}
	// unshift
	//x.rules, x.rules[0] = append(x.rules[:1], x.rules[0:]...), c
}

func (x *sortRules) Clear() {
	x.rules = x.rules[:0]
}

func (x sortRules) Len() int {
	return len(x.t)
}

func (x sortRules) Less(i, j int) bool {
	// from prior key
	for n := len(x.rules) - 1; n >= 0; n-- {
		v := x.rules[n].compare(x.t[i], x.t[j])
		if v != 0 {
			return v < 0
		}
	}

	return true
}

func (x sortRules) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}
