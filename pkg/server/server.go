package server

import (
	"net/http"
    "github.com/gorilla/mux"
)

type Server struct {
	//router *http.ServeMux
	router *mux.Router

	serverMiddleWare []func(http.Handler) http.Handler
}

func NewServer() *Server {
	s := Server{}
	//s.router = http.NewServeMux()
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
	server := NewServer()
	var wrappedServer http.Handler
	wrappedServer = server
	for _, mw := range s.serverMiddleWare {
		wrappedServer2 := mw(wrappedServer)
		wrappedServer = wrappedServer2
	}
	return wrappedServer
}
