package database

import "errors"

const wasteCategoryNuber = 5

type InmemoryDb []map[string]struct{}

func NewInmemoryDb() InmemoryDb {
	db := make(InmemoryDb, wasteCategoryNuber)
	for i := 0; i < wasteCategoryNuber; i++ {
		db[i] = make(map[string]struct{})
	}
	return db
}

func (d InmemoryDb) GetWasteClass(wasteName string) (int, error) {
	for idx, class := range d {
		if _, ok := class[wasteName]; ok {
			return idx + 1, nil
		}
	}
	return 0, ErrNotFound
}

func (d InmemoryDb) SetClassForWaste(wasteClass int, wasteName string) error {
	if wasteClass < 1 || wasteClass > 5 {
		return ErrWrongCategory
	}

	_, err := d.GetWasteClass(wasteName)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			d[wasteClass-1][wasteName] = struct{}{}
		}
	}
	return nil
}

func (d InmemoryDb) GetAll() interface{} {
	return d
}
