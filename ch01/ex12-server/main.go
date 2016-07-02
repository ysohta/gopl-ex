// ex12-server provides http server to display lissajous image.
package main

import (
	"log"
	"net/http"
)

func main() {
	// Query 	Example:
	// http://localhost:8000/?cycles=20&res=0.01&size=200&nframes=32&delay=16
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		p := mapToParam(r.Form)
		lissajous(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
