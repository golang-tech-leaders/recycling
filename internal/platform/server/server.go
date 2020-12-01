package server

import (
	"fmt"
	"net/http"
	"recycling/internal/config"

	"github.com/gorilla/mux"
)

// Server provides the server functionality
type Server struct {
	r       *mux.Router
	db      WasteStorageRepository
	srv     *http.Server
	timeout int
}

// NewServer creates a server and prepares a router
func NewServer(cfg *config.Config, storage WasteStorageRepository) *Server {
	s := Server{
		r:       mux.NewRouter(),
		db:      storage,
		timeout: cfg.ReqTimeoutSec,
	}

	s.setupRouter()

	address := fmt.Sprintf(":%s", cfg.AppPort)
	s.srv = &http.Server{
		Handler: s.r,
		Addr:    address,
	}

	return &s
}

func (s *Server) setupRouter() {
	s.r.HandleFunc("/hello", s.hello).Methods("GET", "POST")
	s.r.HandleFunc("/waste/type/list", s.getWasteTypes).Methods("GET")
	s.r.HandleFunc("/waste/type/search/{text}", s.getWasteTypeByName).Methods("GET")
	s.r.HandleFunc("/waste/type/{type_id}", s.getWasteTypeByID).Methods("GET")

}

// Run starts the server
func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}

// Shutdown closes server
func (s *Server) Shutdown() error {
	return s.srv.Close()
}
