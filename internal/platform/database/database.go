package database

import models "recycling/internal/model"

// WasteStorageRepository interface describes storage contract
type WasteStorage interface {
	GetWasteTypes() ([]models.WasteType, error)
	GetWasteTypeByName(wasteName string) (models.WasteType, error)
	GetWasteTypeByID(wasteTypeID string) (models.WasteType, error)
}
