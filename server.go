package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/api/posts/", home)
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
