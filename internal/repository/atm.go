package repository

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"gorm.io/gorm"
)

type AtmPostgres struct {
	db *gorm.DB
}

func NewAtmPostgres(db *gorm.DB) *AtmPostgres {
	return &AtmPostgres{db: db}
}

func (r *AtmPostgres) Create(atm entity.Atm) (int64, error) {
	result := r.db.Create(&atm)
	if result.Error != nil {
		return 0, fmt.Errorf("AtmPostgres - Create - r.db.Create: %w", result.Error)
	}

	return atm.ID, nil
}

func (r *AtmPostgres) GetAll() ([]entity.Atm, error) {
	var atms []entity.Atm

	result := r.db.Find(&atms)
	if result.Error != nil {
		return nil, fmt.Errorf("AtmPostgres - GetAll - r.db.Find: %w", result.Error)
	}

	return atms, nil
}

func (r *AtmPostgres) GetById(id int64) (entity.Atm, error) {
	var atm entity.Atm

	result := r.db.First(&atm, id)
	if result.Error != nil {
		return atm, fmt.Errorf("AtmPostgres - GetById - r.db.First: %w", result.Error)
	}

	return atm, nil
}

func (r *AtmPostgres) Update(atmId int64, inputAtm entity.Atm) error {
	var atm entity.Atm

	result := r.db.First(&atm, atmId)
	if result.Error != nil {
		return fmt.Errorf("AtmPostgres - Update - r.db.First: %w", result.Error)
	}

	result = r.db.Model(atm).Updates(inputAtm)
	if result.Error != nil {
		return fmt.Errorf("AtmPostgres - Update - r.db.Model().Updates: %w", result.Error)
	}

	return nil
}

func (r *AtmPostgres) Delete(atmId int64) error {
	result := r.db.Delete(&entity.Atm{}, atmId)
	if result.Error != nil {
		return fmt.Errorf("AtmPostgres - Delete - r.db.Delete: %w", result.Error)
	}

	return nil
}