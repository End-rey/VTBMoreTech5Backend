package entity

import (
    _ "gorm.io/gorm"
    "time"
)

type Atm struct {
    ID            int64 `gorm:"primaryKey"`
    Address       string
    Latitude      float32
    Longitude     float32
    WorkTimeStart time.Time `gorm:"type:time;not null"`
    WorkTimeEnd   time.Time `gorm:"type:time;not null"`
	Services []Service `gorm:"many2many:atms_services;foreignKey:AtmID;joinForeignKey:AtmID;References:ID;joinReferences:ServiceID"`
}

func (a *Atm) TableName() string {
    return "atms"
}