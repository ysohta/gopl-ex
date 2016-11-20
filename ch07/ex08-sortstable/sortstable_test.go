package sortstable

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortRules(t *testing.T) {
	tests := []struct {
		rules  []comparer
		values []*Track
		want   []*Track
	}{
		{
			[]comparer{ByTitle, ByArtist},
			[]*Track{
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			},
			[]*Track{
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			},
		},
		{
			[]comparer{ByTitle, ByYear},
			[]*Track{
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			},
			[]*Track{
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			},
		},
		{
			[]comparer{ByArtist, ByTitle},
			[]*Track{
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			},
			[]*Track{
				{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
				{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
				{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
				{"Go", "Moby", "Moby", 1992, length("3m37s")},
			},
		},
	}

	for _, test := range tests {
		sort.Sort(sortRules{test.values, test.rules})

		for i := 0; i < len(test.values); i++ {
			if !reflect.DeepEqual(test.values[i], test.want[i]) {
				t.Errorf("expected:%v actual:%v at %d", test.want[i], test.values[i], i)
			}
		}
	}
}
