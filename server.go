package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Post struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var data = []*Post{
	&Post{"Title 1", "Body 1"},
	&Post{"Title 2", "Body 2"},
	&Post{"Title 3", "Body 3"},
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/api/posts/", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
