package main

import (
	"net/http"
	"time"
	"encoding/json"
	"strconv"
	"errors"
)

func main() {
	http.HandleFunc("/", todoIndex)
	http.HandleFunc("/todo", todo)

	http.ListenAndServe(":8080",nil)
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	println("todoIndex")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Todos)
}

func todo(w http.ResponseWriter, r *http.Request) {
	println("todo")
	idStr:=r.URL.Query().Get("id")
	id, err:=strconv.Atoi(idStr)
	if err!=nil{
		panic(errors.New("invalid id"))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err:=json.NewEncoder(w).Encode(Todos[id]); err!=nil{
		panic(err)
	}
}

type Todo struct {
	Title string
	Due   time.Time
}

var Todos = []Todo{
	Todo{"Do this", time.Now().Add(3 * time.Hour)},
	Todo{"Do that", time.Now().Add(10 * time.Hour)},
}
