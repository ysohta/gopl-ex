package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.del)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid item name: %q\n", item)
		return
	}
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")

	var f float64
	var err error
	if f, err = strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	db[item] = dollars(f)
	fmt.Fprintf(w, "created %s: %s\n", item, db[item])

}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid item name: %q\n", item)
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item not found: %q\n", item)
		return
	}

	price := req.URL.Query().Get("price")

	var f float64
	var err error
	if f, err = strconv.ParseFloat(price, 32); err != nil {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}
	db[item] = dollars(f)
	fmt.Fprintf(w, "updated %s: %s\n", item, db[item])
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "invalid item name: %q\n", item)
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item not found: %q\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "item deleted %s\n", item)
}
