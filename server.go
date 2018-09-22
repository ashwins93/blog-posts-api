package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/posts", allPosts)
	http.HandleFunc("/api/posts/", onePost)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
