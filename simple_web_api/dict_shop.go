package main

import (
	"fmt"
	"net/http"
)

func main() {
	db := database{"book1": 10, "book2": 20}

	muxer := http.NewServeMux()
	muxer.Handle("/list", http.HandlerFunc(db.list))
	muxer.Handle("/price", http.HandlerFunc(db.price))
	http.ListenAndServe(":8080", muxer)
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for key, value := range db {
		fmt.Fprintf(w, "item: %s at: %s\n", key, value)
	}
}
func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok{
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprintf(w, "Item %s not found", item)

		return
	}
	fmt.Fprintf(w,"price for %s is %s", item, price)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for key, value := range db {
			fmt.Fprintf(w, "item: %s at: %s\n", key, value)
		}
	case "/item":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok{
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintf(w, "Item %s not found", item)

			return
		}
		fmt.Fprintf(w,"price for %s is %s", item, price)
	default:
		http.NotFound(w, r)
	}
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%f", d)
}
