package repository

import "gorm.io/gorm"

type AlgorithmPostgres struct {
	db *gorm.DB
}

func NewAlgorithmPostgres(db *gorm.DB) *AlgorithmPostgres {
	return &AlgorithmPostgres{db: db}
}