package models

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

// NullString extends sql.NullString
type NullString sql.NullString

// WasteType provides type of waste
type WasteType struct {
	ID          NullString `json:"id"`
	Name        NullString `json:"name"`
	Description NullString `json:"description"`
}

// WasteTypeList is a slice of WasteType structs
type WasteTypeList []WasteType

// MarshalJSON provides json.Marshal() over sql.NullString.String
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if ns.Valid {
		return json.Marshal(ns.String)
	}
	return []byte(`null`), nil
}

// Scan scans sql.NullString
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*ns = NullString{s.String, false}
	} else {
		*ns = NullString{s.String, true}
	}

	return nil
}
