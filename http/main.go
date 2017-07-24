package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/root", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root file"))
	})
	mux.HandleFunc("/root/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root folder"))
	})

	http.ListenAndServe(":9090", mux)
}
