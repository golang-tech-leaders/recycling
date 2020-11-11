package database

import (
	"database/sql"
	"fmt"
	"log"
	models "recycling/internal/model"

	_ "github.com/lib/pq" // required for PostgreSQL connection
)

// PostgresWasteStorage incapsulates PostgreSQL storage
type PostgresWasteStorage struct {
	db *sql.DB
}

// NewPostgresWasteStorage creates and returns an instance of PostgresWasteStorage
func NewPostgresWasteStorage(config *models.Config) *PostgresWasteStorage {
	address := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err := sql.Open("postgres", address)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	pgStorage := PostgresWasteStorage{db: db}
	return &pgStorage
}

// GetWasteTypes returns a list of all available waste types
func (p *PostgresWasteStorage) GetWasteTypes() (models.WasteTypeList, error) {
	rows, err := p.db.Query("SELECT * FROM waste_type ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wasteTypes := make(models.WasteTypeList, 0)
	for rows.Next() {
		var wt models.WasteType
		if err := rows.Scan(&wt.ID, &wt.Name, &wt.Description); err != nil {
			return nil, err
		}
		wasteTypes = append(wasteTypes, wt)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return wasteTypes, nil
}

// GetWasteTypeByName returns WasteType by it's name
func (p *PostgresWasteStorage) GetWasteTypeByName(wasteName string) (models.WasteType, error) {
	var wt models.WasteType
	err := p.db.QueryRow(`SELECT id, name, description FROM waste_type WHERE name like $1;`, "%"+wasteName+"%").Scan(&wt.ID, &wt.Name, &wt.Description)
	switch err {
	case sql.ErrNoRows:
		return models.WasteType{}, ErrNotFound
	case nil:
		return wt, nil
	default:
		return models.WasteType{}, err
	}
}

// GetWasteTypeByID returns WasteType by ID
func (p *PostgresWasteStorage) GetWasteTypeByID(wasteTypeID string) (models.WasteType, error) {
	var wt models.WasteType
	err := p.db.QueryRow(`SELECT id, name, description FROM waste_type WHERE id = $1;`, wasteTypeID).Scan(&wt.ID, &wt.Name, &wt.Description)
	switch err {
	case sql.ErrNoRows:
		return models.WasteType{}, ErrNotFound
	case nil:
		return wt, nil
	default:
		return models.WasteType{}, err
	}
}
