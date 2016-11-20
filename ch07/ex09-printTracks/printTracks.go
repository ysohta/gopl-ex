package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

var trackList = template.Must(template.New("trackList").
	Parse(`
<h1>Tracks</h1>
<button type='button' onclick="window.location = window.location.pathname">Reset</button>
<table>
<tr style='text-align: left'>
  <th><a onclick="window.location+=((window.location.href.indexOf('?')+1)?'':'?')+'&sortby=title'">Title</a></th>
  <th><a onclick="window.location+=((window.location.href.indexOf('?')+1)?'':'?')+'&sortby=artist'">Artist</a></th>
  <th><a onclick="window.location+=((window.location.href.indexOf('?')+1)?'':'?')+'&sortby=album'">Album</a></th>
  <th><a onclick="window.location+=((window.location.href.indexOf('?')+1)?'':'?')+'&sortby=year'">Year</a></th>
  <th><a onclick="window.location+=((window.location.href.indexOf('?')+1)?'':'?')+'&sortby=len'">Length</a></th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

const (
	file = "index.html"
)

var (
	myTracks sortRules = sortRules{tracks, []comparer{}}
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		val, ok := r.Form["sortby"]
		if ok {

			log.Println(val)
			for _, term := range val {
				switch term {
				case "title":
					myTracks.Add(ByTitle)
				case "artist":
					myTracks.Add(ByArtist)
				case "album":
					myTracks.Add(ByAlbum)
				case "year":
					myTracks.Add(ByYear)
				case "len":
					myTracks.Add(ByLength)
				default:
					log.Println("invalid term:", term)
				}
			}
			sort.Sort(myTracks)
		} else {
			myTracks.Clear()
		}

		if err := trackList.Execute(w, tracks); err != nil {
			log.Fatal(err)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
