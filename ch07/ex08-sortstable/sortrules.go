package sortstable

var (
	ByTitle  comparer = &byTitle{}
	ByArtist comparer = &byArtist{}
	ByAlbum  comparer = &byAlbum{}
	ByYear   comparer = &byYear{}
	ByLength comparer = &byLength{}
)

type byTitle []*Track

func (b *byTitle) compare(t1, t2 *Track) int {
	if t1.Title == t2.Title {
		return 0
	}

	if t1.Title < t2.Title {
		return -1
	}
	return 1
}

func (x byTitle) Len() int { return len(x) }

func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }

func (x byTitle) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type byArtist []*Track

func (b *byArtist) compare(t1, t2 *Track) int {
	if t1.Artist == t2.Artist {
		return 0
	}

	if t1.Artist < t2.Artist {
		return -1
	}
	return 1
}

func (x byArtist) Len() int { return len(x) }

func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }

func (x byArtist) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type byAlbum []*Track

func (b *byAlbum) compare(t1, t2 *Track) int {
	if t1.Album == t2.Album {
		return 0
	}

	if t1.Album < t2.Album {
		return -1
	}
	return 1
}

func (x byAlbum) Len() int { return len(x) }

func (x byAlbum) Less(i, j int) bool { return x[i].Album < x[j].Album }

func (x byAlbum) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (b *byYear) compare(t1, t2 *Track) int {
	if t1.Year == t2.Year {
		return 0
	}

	if t1.Year < t2.Year {
		return -1
	}
	return 1
}

func (x byYear) Len() int { return len(x) }

func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }

func (x byYear) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type byLength []*Track

func (b *byLength) compare(t1, t2 *Track) int {
	if t1.Length == t2.Length {
		return 0
	}

	if t1.Length < t2.Length {
		return -1
	}
	return 1
}

func (x byLength) Len() int { return len(x) }

func (x byLength) Less(i, j int) bool { return x[i].Length < x[j].Length }

func (x byLength) Swap(i, j int) { x[i], x[j] = x[j], x[i] }
