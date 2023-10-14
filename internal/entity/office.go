package entity

import (
	_ "gorm.io/gorm"
)

type Office struct {
	ID            int64     `gorm:"primaryKey"`
	SalePointName string    `gorm:"type:varchar(150)"`
	Address       string    `gorm:"type:varchar(150)"`
	Status        bool
	OfficeType    string    `gorm:"type:varchar(100)"`
	Latitude      float32   `gorm:"not null"`
	Longitude     float32   `gorm:"not null"`
	MetroStation  string    `gorm:"type:varchar(150)"`
	Services      []Service `gorm:"many2many:offices_services"`
	Charts        []ChartsOffices
	WorksTime     []WorksTimeOffice
}