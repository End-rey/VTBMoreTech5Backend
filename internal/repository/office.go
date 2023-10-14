package repository

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"gorm.io/gorm"
)

type OfficePostgres struct {
	db *gorm.DB
}

func NewOfficePostgres(db *gorm.DB) *OfficePostgres {
	return &OfficePostgres{db: db}
}

func (r *OfficePostgres) Create(office entity.Office) (int64, error) {
	result := r.db.Create(&office)
	if result.Error != nil {
		return 0, fmt.Errorf("OfficePostgres - Create - r.db.Create: %w", result.Error)
	}

	return office.ID, nil
}

func (r *OfficePostgres) GetAll() ([]entity.Office, error) {
	var offices []entity.Office

	result := r.db.Find(&offices)
	if result.Error != nil {
		return nil, fmt.Errorf("OfficePostgres - GetAll - r.db.Find: %w", result.Error)
	}
	
	fmt.Printf("offices %+v", offices)

	return offices, nil
}

func (r *OfficePostgres) GetById(officeId int64) (entity.Office, error) {
	var office entity.Office

	result := r.db.First(&office, officeId)
	if result.Error != nil {
		return office, fmt.Errorf("OfficePostgres - GetById - r.db.First: %w", result.Error)
	}

	return office, nil
}

func (r *OfficePostgres) Update(officeId int64, inputOffice entity.Office) error {
	var office entity.Office

	result := r.db.First(&office, officeId)
	if result.Error != nil {
		return fmt.Errorf("OfficePostgres - Update - r.db.First: %w", result.Error)
	}

	result = r.db.Model(office).Updates(inputOffice)
	if result.Error != nil {
		return fmt.Errorf("OfficePostgres - Update - r.db.Model().Updates: %w", result.Error)
	}

	return nil
}

func (r *OfficePostgres) Delete(officeId int64) error {
	result := r.db.Delete(&entity.Office{}, officeId)
	if result.Error != nil {
		return fmt.Errorf("OfficePostgres - Delete - r.db.Delete: %w", result.Error)
	}

	return nil
}
