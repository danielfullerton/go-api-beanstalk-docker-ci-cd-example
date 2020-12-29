package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// todo: add your route handlers here...
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	log.Println("server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
