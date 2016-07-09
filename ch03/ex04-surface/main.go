// ex04-surface provides http server to display SVG rendering of a 3-D surface function.
package main

import (
	"log"
	"net/http"
)

func main() {
	// Query 	Example:
	// http://localhost:8000/?width=1000&height=600&topColor=00ff00ff&bottomColor=ffffffff
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		p := mapToParam(r.Form)
		surface(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
