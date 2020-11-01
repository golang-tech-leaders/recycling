package models

import "database/sql"

// WasteType provides type of waste
type WasteType struct {
	ID          sql.NullString `json:"id"`
	Name        sql.NullString `json:"name"`
	Description sql.NullString `json:"description"`
}

// WasteTypeList is a slice of WasteType structs
type WasteTypeList []WasteType
