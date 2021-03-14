package server

import (
	"fmt"
	"net/http"
)

func (s *Server) handleAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// do something with it.
		fmt.Fprintf(w, "blah blah blah")
	}
}
