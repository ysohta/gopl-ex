package sortstable

var (
	ByTitle  comparer = &byTitle{}
	ByArtist comparer = &byArtist{}
	ByAlbum  comparer = &byAlbum{}
	ByYear   comparer = &byYear{}
	ByLength comparer = &byLength{}
)

type byTitle struct{}

func (b *byTitle) compare(t1, t2 *Track) int {
	if t1.Title == t2.Title {
		return 0
	}

	if t1.Title < t2.Title {
		return -1
	}
	return 1
}

type byArtist struct{}

func (b *byArtist) compare(t1, t2 *Track) int {
	if t1.Artist == t2.Artist {
		return 0
	}

	if t1.Artist < t2.Artist {
		return -1
	}
	return 1
}

type byAlbum struct{}

func (b *byAlbum) compare(t1, t2 *Track) int {
	if t1.Album == t2.Album {
		return 0
	}

	if t1.Album < t2.Album {
		return -1
	}
	return 1
}

type byYear struct{}

func (b *byYear) compare(t1, t2 *Track) int {
	if t1.Year == t2.Year {
		return 0
	}

	if t1.Year < t2.Year {
		return -1
	}
	return 1
}

type byLength struct{}

func (b *byLength) compare(t1, t2 *Track) int {
	if t1.Length == t2.Length {
		return 0
	}

	if t1.Length < t2.Length {
		return -1
	}
	return 1
}
