package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	s := NewServer()

	http.ListenAndServe(":3333", s.Handler())
}

type Server struct {
	router http.Handler
}

func NewServer() *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	s := &Server{
		router: r,
	}

	r.Get("/", s.HandleIndex)

	return s
}

func (s *Server) Handler() http.Handler {
	return s.router
}

func (s *Server) HandleIndex(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("hello world"))
}
