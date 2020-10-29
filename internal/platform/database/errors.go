package database

import "errors"

var (
	// ErrWrongCategory waste category number is wrong
	ErrWrongCategory = errors.New("Category number should be more than 0 and less than 6")
	// ErrNotFound not found
	ErrNotFound = errors.New("Waste not found")
)
