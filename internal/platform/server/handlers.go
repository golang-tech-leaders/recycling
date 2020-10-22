package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"recycling/internal/platform/database"
	"recycling/internal/waste"

	"github.com/gorilla/mux"
)

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
}

func (s *Server) showAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%+v", s.db.GetAll())
}

func (s *Server) getWasteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wasteName := vars["waste_name"]

	wasteClass, err := s.db.GetWasteClass(wasteName)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No waste class was found for waste: "+wasteName, http.StatusNotFound)
			return
		}
	}
	fmt.Fprintf(w, "Waste %q is in the class %d", wasteName, wasteClass)
}

func (s *Server) newWaste(w http.ResponseWriter, r *http.Request) {
	var req waste.Waste

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Unable to store due to: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.SetClassForWaste(req.Class, req.Name)
	if err != nil {
		if errors.Is(err, database.ErrWrongCategory) {
			http.Error(w, "Wrong waste category provided: "+err.Error(), http.StatusNotModified)
			return
		}
		http.Error(w, "Something went wrong"+err.Error(), http.StatusNotImplemented)
	}
}
