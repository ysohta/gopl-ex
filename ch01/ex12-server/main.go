// ex12-server provides http server to display lissajous image.
package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Query 	Example:
	// http://localhost:8000/?cycles=20&res=0.01&size=200&nframes=32&delay=16
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := param{
			cycles:  parseFormValueToInt(r, "cycles", 5),
			res:     parseFormValueToFloat64(r, "res", 0.001),
			size:    parseFormValueToInt(r, "size", 100),
			nframes: parseFormValueToInt(r, "nframes", 64),
			delay:   parseFormValueToInt(r, "delay", 8),
		}
		lissajous(w, p)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// parseFormValueToInt returns int value parsed from request form.
// defaultVal is returned when parsing is failed.
func parseFormValueToInt(r *http.Request, key string, defaultVal int) int {
	v := r.FormValue(key)
	if v == "" {
		return defaultVal
	}
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	return defaultVal
}

// parseFormValueToInt returns flaot64 value parsed from request form.
// defaultVal is returned when parsing is failed.
func parseFormValueToFloat64(r *http.Request, key string, defaultVal float64) float64 {
	v := r.FormValue(key)
	if v == "" {
		return defaultVal
	}
	if f, err := strconv.ParseFloat(v, 64); err == nil {
		return f
	}
	return defaultVal
}
