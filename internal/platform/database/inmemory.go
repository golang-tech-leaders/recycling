package database

import "errors"

type InmemoryDb map[int][]string

func NewInmemoryDb() InmemoryDb {
	return make(InmemoryDb)
}

func (d InmemoryDb) GetWasteClass(s string) (int, error) {
	for wasteClass, wasteNames := range d {
		for _, n := range wasteNames {
			if n == s {
				return wasteClass, nil
			}
		}
	}
	return -1, errors.New("Not found")
}

func (d InmemoryDb) SetClassForWaste(i int, s string) error {
	if i < 1 || i > 5 {
		return errors.New("Waste class should be >= 1 and <= 5")
	}
	cls, _ := d.GetWasteClass(s)
	if cls == -1 {
		waste := d[i]
		waste = append(waste, s)
		d[i] = waste
	}
	return nil
}
