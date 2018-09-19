package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var data = []*Post{
	&Post{1, "Title 1", "Body 1"},
	&Post{2, "Title 2", "Body 2"},
	&Post{3, "Title 3", "Body 3"},
}

var currentID = 4

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		err := enc.Encode(data)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if r.Method == "POST" {
		r.ParseForm()
		var p Post
		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&p); err != nil {
			log.Fatal(err)
		}
		p.ID = currentID
		currentID++
		data = append(data, &p)
		enc := json.NewEncoder(w)

		w.WriteHeader(http.StatusCreated)
		if err := enc.Encode(p); err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	http.HandleFunc("/api/posts", home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
