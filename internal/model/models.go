package models

// WasteType provides type of waste
type WasteType struct {
	TypeID string `json:"type_id"`
	Name   string `json:"type_name"`
}

// Waste provides waste name and waste type
type Waste struct {
	ID   int    `json:"waste_id"`
	Name string `json:"waste_name"`
	Type WasteType
}
