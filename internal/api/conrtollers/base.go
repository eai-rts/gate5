package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server ...
type Server struct {
	Router *mux.Router
}

// Run the server
func (s *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, s.Router))
}

// NewServer creates a new server instance
func NewServer() *Server {
	return &Server{Router: mux.NewRouter()}
}
