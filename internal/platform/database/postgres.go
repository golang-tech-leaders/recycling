package database

import (
	"database/sql"
	"log"
	models "recycling/internal/model"

	_ "github.com/lib/pq"
)

type PostgresWasteStorage struct {
	db *sql.DB
}

func NewPostgresWasteStorage(address string) *PostgresWasteStorage {
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

func (p *PostgresWasteStorage) GetWasteTypes() ([]models.WasteType, error) {
	rows, err := p.db.Query("SELECT * FROM waste_type ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	wasteTypes := make([]models.WasteType, 0)
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

func (p *PostgresWasteStorage) GetWasteTypeByName(wasteName string) (models.WasteType, error) {
	return models.WasteType{}, nil
}

func (p *PostgresWasteStorage) GetWasteTypeByID(wasteTypeID string) (models.WasteType, error) {
	return models.WasteType{}, nil
}

func (p *PostgresWasteStorage) PopulateWasteTypes() {
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type1', 'waste type 1')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type2', 'waste type 2')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type3', 'waste type 3')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type4', 'waste type 4')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type5', 'waste type 5')")
	p.db.Exec("INSERT INTO waste_type(id, name, description) VALUES ('type6', 'waste type 6', 'description 6')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type7', 'waste type 7')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type8', 'waste type 8')")
	p.db.Exec("INSERT INTO waste_type(id, name) VALUES ('type9', 'waste type 9')")
}
