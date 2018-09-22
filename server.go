package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

var idGrabber = regexp.MustCompile("^/api/posts/(\\d+)$")

func allPosts(w http.ResponseWriter, r *http.Request) {
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

func onePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := idGrabber.FindStringSubmatch(r.URL.Path)
	if len(id) == 0 {
		http.NotFound(w, r)
		return
	}
	parsed, err := strconv.ParseInt(id[1], 10, 32)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	if r.Method == "GET" {
		for _, v := range data {
			if parsed == int64(v.ID) {
				enc := json.NewEncoder(w)
				err = enc.Encode(v)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}
		http.NotFound(w, r)
		return
	}

}

func main() {
	http.HandleFunc("/api/posts", allPosts)
	http.HandleFunc("/api/posts/", onePost)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
