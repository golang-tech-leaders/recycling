package server

import (
	"context"
	models "recycling/internal/model"
)

// WasteStorageRepository interface describes storage contract
type WasteStorageRepository interface {
	GetWasteTypes(ctx context.Context) (models.WasteTypeList, error)
	GetWasteTypeByName(ctx context.Context, wasteName string) (models.WasteType, error)
	GetWasteTypeByID(ctx context.Context, wasteTypeID string) (models.WasteType, error)
}
