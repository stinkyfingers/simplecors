package simplecors

import "net/http"

// This is a package-level example:
func Example() {
	mux := http.NewServeMux()
	cors := DefaultCors()
	mux.HandleFunc("/", cors.Middleware(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success!"))
	}))
	http.ListenAndServe(":8080", mux)
}
