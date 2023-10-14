package entity

import (
	_ "gorm.io/gorm"
)

type Service struct {
	ID          int64    `gorm:"primaryKey"`
	ServiceName string   `gorm:"type:varchar(250);not null"`
	Atms        []Atm    `gorm:"many2many:atms_services;foreignKey:ServiceID;joinForeignKey:ServiceID;References:ID;joinReferences:AtmID"`
	Offices     []Office `gorm:"many2many:offices_services;foreignKey:ServiceID;joinForeignKey:ServiceID;References:ID;joinReferences:OfficeID"`
}

func (s *Service) TableName() string {
	return "services"
}
