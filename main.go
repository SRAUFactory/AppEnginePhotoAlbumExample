package main

import "fmt"
import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Action")
}

func photos(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Photos Action")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete Action")
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post Action")
}

func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/photos", photos)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/post", post)
}
