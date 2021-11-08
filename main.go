package main

import (
	"clipboard/storage"
	"fmt"
	"io"
	"net/http"
)

var store = storage.NewStore()

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/get", get)
	err := http.ListenAndServe("localhost:80", nil)
	if err != nil {
		fmt.Println("fail to start server")
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	content := query.Get("content")
	key := store.Save(content)
	io.WriteString(w, key)
}

func get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	key := query.Get("id")
	content := store.Get(key)
	io.WriteString(w, content)
}
