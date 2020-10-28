package server

import (
	"errors"
	"fmt"
	"net/http"

	"recycling/internal/platform/database"

	"github.com/gorilla/mux"
)

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
}

func (s *Server) getWasteTypes(w http.ResponseWriter, r *http.Request) {
	wasteTypes, err := s.db.GetWasteTypes()
	if err != nil {
		http.Error(w, "Unable to get list of waste types", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "%+v", wasteTypes)
}

func (s *Server) getTypeByWasteName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wasteName := vars["name"]
	wasteType, err := s.db.GetTypeByWaste(wasteName)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No waste type was found for waste: "+wasteName, http.StatusNotFound)
			return
		}
	}
	fmt.Fprintf(w, "%+v", wasteType)
}

func (s *Server) getWasteByTypeID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wasteTypeID := vars["id"]
	wasteList, err := s.db.GetWasteByType(wasteTypeID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No waste found for type: "+wasteTypeID, http.StatusNotFound)
			return
		}
	}
	fmt.Fprintf(w, "%+v", wasteList)
}
