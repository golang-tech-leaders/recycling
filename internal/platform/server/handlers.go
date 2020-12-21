package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"recycling/internal/platform/database"

	"github.com/gorilla/mux"
)

func (s *Server) hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, request: %+v", r)
}

func (s *Server) getWasteTypes(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug("getWasteTypes: canceling context")
		cancel()
	}()
	wasteTypes, err := s.db.GetWasteTypes(ctx)
	if err != nil {
		http.Error(w, "Unable to get list of waste types due to: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&wasteTypes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getWasteTypeByName(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug("getWasteTypeByName: canceling context")
		cancel()
	}()
	vars := mux.Vars(r)
	wasteName := vars["text"]
	wasteType, err := s.db.GetWasteTypeByName(ctx, wasteName)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No waste type was found for waste: "+wasteName, http.StatusNotFound)
			return
		}
	}

	err = json.NewEncoder(w).Encode(&wasteType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *Server) getWasteTypeByID(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(s.timeout)*time.Second)
	defer func() {
		s.log.Debug("getWasteTypeByID: canceling context")
		cancel()
	}()
	vars := mux.Vars(r)
	wasteTypeID := vars["type_id"]
	wasteType, err := s.db.GetWasteTypeByID(ctx, wasteTypeID)
	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			http.Error(w, "No waste found for type: "+wasteTypeID, http.StatusNotFound)
			return
		}
	}

	err = json.NewEncoder(w).Encode(&wasteType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
