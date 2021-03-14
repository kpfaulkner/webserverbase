package server

import (
	"net/http"
)

func (s *Server) isAdmin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// check if current user is admin (somehow?)
		h(w, r)
	}
}
