package database

import (
	"context"
	"database/sql"
	"fmt"

	"recycling/internal/config"
	"recycling/internal/logger"
	models "recycling/internal/model"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // required for go-migrate via files
	_ "github.com/lib/pq"                                // required for PostgreSQL connection
)

// PostgresWasteStorage incapsulates PostgreSQL storage
type PostgresWasteStorage struct {
	db  *sql.DB
	log *logger.Logger
}

// NewPostgresWasteStorage creates and returns an instance of PostgresWasteStorage
func NewPostgresWasteStorage(config *config.DBConfig, logger *logger.Logger) *PostgresWasteStorage {
	dbURL := config.DbURL
	if dbURL == "" {
		dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	}
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		logger.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logger.Fatal(err)
	}

	pgStorage := PostgresWasteStorage{db: db, log: logger}
	return &pgStorage
}

// Migrate ups version of DB model
func (p *PostgresWasteStorage) Migrate() {
	driver, err := postgres.WithInstance(p.db, &postgres.Config{})
	if err != nil {
		p.log.Fatal("[MIGRATE] Unable to get driver due to: " + err.Error())
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		//"file:///Users/daniel.khasanov/techleaders/recycling/migrations",
		"postgres", driver)
	if err != nil {
		p.log.Fatal("[MIGRATE] Unable to get migrate instance due to: " + err.Error())
	}
	err = m.Up()
	switch err {
	case migrate.ErrNoChange:
		return
	case nil:
		return
	default:
		p.log.Fatal("[MIGRATE] Unable to apply DB migrations due to: " + err.Error())
	}
}

// GetWasteTypes returns a list of all available waste types
func (p *PostgresWasteStorage) GetWasteTypes(ctx context.Context) (models.WasteTypeList, error) {
	rows, err := p.db.QueryContext(ctx, "SELECT * FROM waste_type ORDER BY id")
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
func (p *PostgresWasteStorage) GetWasteTypeByName(ctx context.Context, wasteName string) (*models.WasteType, error) {
	var wt models.WasteType
	err := p.db.QueryRowContext(ctx, `SELECT id, name, description FROM waste_type WHERE name like $1;`, "%"+wasteName+"%").Scan(&wt.ID, &wt.Name, &wt.Description)
	switch err {
	case sql.ErrNoRows:
		return nil, ErrNotFound
	case nil:
		return &wt, nil
	default:
		return nil, err
	}
}

// GetWasteTypeByID returns WasteType by ID
func (p *PostgresWasteStorage) GetWasteTypeByID(ctx context.Context, wasteTypeID string) (*models.WasteType, error) {
	var wt models.WasteType
	err := p.db.QueryRowContext(ctx, `SELECT id, name, description FROM waste_type WHERE id = $1;`, wasteTypeID).Scan(&wt.ID, &wt.Name, &wt.Description)
	switch err {
	case sql.ErrNoRows:
		return nil, ErrNotFound
	case nil:
		return &wt, nil
	default:
		return nil, err
	}
}
