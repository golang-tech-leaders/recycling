package server

import (
	"net/http"
	"recycling/internal/platform/database"

	"github.com/gorilla/mux"
)

// Server provides the server functionality
type Server struct {
	r   *mux.Router
	db  database.WasteStorage
	srv *http.Server
}

// NewServer creates a server and prepares a router
func NewServer(address string, storage database.WasteStorage) *Server {
	s := Server{
		r:  mux.NewRouter(),
		db: storage,
	}

	s.setupRouter()

	s.srv = &http.Server{
		Handler: s.r,
		Addr:    address,
	}

	return &s
}

func (s *Server) setupRouter() {
	s.r.HandleFunc("/hello", s.hello).Methods("GET", "POST")
	s.r.HandleFunc("/api/waste/type/list", s.getWasteTypes).Methods("GET")
	s.r.HandleFunc("/api/waste/type/{name}", s.getTypeByWasteName).Methods("GET")
	s.r.HandleFunc("/api/waste/type/{id}", s.getWasteByTypeID).Methods("GET")
}

// Run starts the server
func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

// Shutdown closes server
func (s *Server) Shutdown() error {
	return s.srv.Close()
}
