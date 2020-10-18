package server

import (
	"net/http"
	"recycling/internal/platform/database"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	r    *mux.Router
	db   database.WasteStorage
}

func NewServer(address string, storage database.WasteStorage) *Server {
	s := Server{
		addr: address,
		r:    mux.NewRouter(),
		db:   storage,
	}
	s.setupRouter()
	return &s
}

func (s *Server) setupRouter() {
	s.r.HandleFunc("/hello", s.hello).Methods("GET", "POST")
	s.r.HandleFunc("/get_class/{waste_name:[A-Za-z]+}", s.getWasteClass).Methods("GET")
	s.r.HandleFunc("/add", s.newWaste).Methods("POST")
}

func (s *Server) Run() error {
	srv := &http.Server{
		Handler: s.r,
		Addr:    s.addr,
	}

	return srv.ListenAndServe()
}
