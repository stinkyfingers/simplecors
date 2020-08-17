// SimpleCors provides a quick, simple means to implement Cors into http servers that
// do not utilize packages like Gorilla with more sophsiticated routing and Cors management.
package simplecors

import (
	"net/http"
)

// Cors is a struct that contains configuration fields.
type Cors struct {
	AllowMethods string
	AllowOrigin  string
	AllowHeaders string
}

// DefaultCors returns a Cors object with field populated by standard defaults.
func DefaultCors() *Cors {
	return &Cors{
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowOrigin:  "*",
		AllowHeaders: "Content-Type",
	}
}

// Middleware accepts and returns a HandlerFunc, handling Cors requests.
func (c *Cors) Middleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", c.AllowMethods)
		w.Header().Set("Access-Control-Allow-Origin", c.AllowOrigin)
		w.Header().Set("Access-Control-Allow-Headers", c.AllowHeaders)
		if r.Method == "OPTIONS" {
			return
		}
		fn(w, r)
	}
}
