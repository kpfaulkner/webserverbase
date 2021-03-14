package server

import (
	"net/http"
    "github.com/gorilla/mux"
)

type Server struct {

	// key used for JWT validation.
	jwtKey string
	router *mux.Router

	serverMiddleWare []func(http.Handler) http.Handler
}

func NewServer(jwtKey string) *Server {
	s := Server{}
	s.jwtKey = jwtKey
	s.router = mux.NewRouter()
	s.routes()

	return &s
}

/*
func NewServerWithMiddleware(mws []func(http.Handler) http.Handler) http.Handler {
	server := NewServer()

	var wrappedServer http.Handler
	wrappedServer = server
	for _, mw := range mws {
		wrappedServer2 := mw(wrappedServer)
		wrappedServer = wrappedServer2
	}
	return wrappedServer
}
*/

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// Use adds middleware.
func (s *Server) Use(middleware func(http.Handler) http.Handler) error {
	s.serverMiddleWare = append(s.serverMiddleWare, middleware)
	return nil
}

// GetServerWithMiddleware returns the server with middleware applied.
func (s *Server) GetServerWithMiddleware() http.Handler {
	server := NewServer(s.jwtKey)
	var wrappedServer http.Handler
	wrappedServer = server
	for _, mw := range s.serverMiddleWare {
		wrappedServer2 := mw(wrappedServer)
		wrappedServer = wrappedServer2
	}
	return wrappedServer
}
