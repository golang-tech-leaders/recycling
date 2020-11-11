package server

import models "recycling/internal/model"

// WasteStorageRepository interface describes storage contract
type WasteStorageRepository interface {
	GetWasteTypes() (models.WasteTypeList, error)
	GetWasteTypeByName(wasteName string) (models.WasteType, error)
	GetWasteTypeByID(wasteTypeID string) (models.WasteType, error)
}
