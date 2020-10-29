package database

import models "recycling/internal/model"

// WasteStorage interface describes storage contract
type WasteStorage interface {
	GetWasteTypes() ([]models.WasteType, error)
	GetTypeByWaste(wasteName string) (models.WasteType, error)
	GetWasteTypeByType(wasteTypeID string) (models.WasteType, error)
}
