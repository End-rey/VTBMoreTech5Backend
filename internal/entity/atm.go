package entity

import (
	"time"

	_ "gorm.io/gorm"
)

type Atm struct {
	ID             int64     `gorm:"primaryKey"`
	Address        string
	Latitude       float32   `gorm:"not null"`
	Longitude      float32   `gorm:"not null"`
	WorkTimeStart  time.Time `gorm:"type:time;not null"`
	WorkTimeEnd    time.Time `gorm:"type:time;not null"`
	Services       []Service `gorm:"many2many:atms_services"`
}