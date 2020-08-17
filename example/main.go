package main

import (
	"net/http"

	"github.com/stinkyfingers/simplecors"
)

func main() {
	mux := http.NewServeMux()
	cors := simplecors.DefaultCors()
	mux.HandleFunc("/", cors.Middleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success!"))
	}))
	http.ListenAndServe(":8080", mux)
}
