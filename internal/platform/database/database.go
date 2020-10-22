package database

// WasteStorage interface describes storage contract
type WasteStorage interface {
	GetWasteClass(s string) (int, error)
	SetClassForWaste(i int, s string) error
	GetAll() interface{}
}
