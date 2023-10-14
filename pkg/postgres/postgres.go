// Package postgres implements postgres connection.
package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New -.
func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
	    return nil, fmt.Errorf("postgres - NewPostgres - gorm.Open: %w", err)
	}

	return db, nil
}
