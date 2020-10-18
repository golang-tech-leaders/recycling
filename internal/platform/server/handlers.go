package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"recycling/internal/waste"

	"github.com/gorilla/mux"
)

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
}

func (s *Server) getWasteClass(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wasteName := vars["waste_name"]
	wasteClass, _ := s.db.GetWasteClass(wasteName)

	if wasteClass == -1 {
		http.Error(w, "not found", http.StatusNotFound)
		return
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
		http.Error(w, "Unable to store due to: "+err.Error(), http.StatusBadRequest)
		return
	}
}
