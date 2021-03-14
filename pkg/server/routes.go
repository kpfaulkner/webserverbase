package server

// Handle ALL routes for the server.
// This will be covering all the end points... and wont have
// separate structs or methods off structs handling registration. Do it all
// here so we just have one place to look!
func (s *Server) routes() {
	s.router.HandleFunc("/api/", s.checkJWT(s.isAdmin(s.handleAPI()))).Methods("GET")
	s.router.HandleFunc("/greetings", s.handleGreeting("hello"))
	s.router.HandleFunc("/greetings/{name}", s.handleGreeting("hello")).Methods("GET")
	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
}
