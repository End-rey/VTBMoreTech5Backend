package repository

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"gorm.io/gorm"
)

type ServicePostgres struct {
	db *gorm.DB
}

func NewServicePostgres(db *gorm.DB) *ServicePostgres {
	return &ServicePostgres{db: db}
}

func (r *ServicePostgres) Create(service entity.Service) (int64, error) {
	result := r.db.Create(&service)
	if result.Error != nil {
		return 0, fmt.Errorf("ServicePostgres - Create - r.db.Create: %w", result.Error)
	}

	return service.ID, nil
}

func (r *ServicePostgres) GetAll() ([]entity.Service, error) {
	var services []entity.Service

	result := r.db.Find(&services)
	if result.Error != nil {
		return nil, fmt.Errorf("ServicePostgres - GetAll - r.db.Find: %w", result.Error)
	}

	return services, nil
}

func (r *ServicePostgres) GetById(serviceId int64) (entity.Service, error) {
	var service entity.Service

	result := r.db.First(&service, serviceId)
	if result.Error != nil {
		return service, fmt.Errorf("ServicePostgres - GetById - r.db.First: %w", result.Error)
	}

	return service, nil
}

func (r *ServicePostgres) Update(serviceId int64, inputOffice entity.Service) error {
	var service entity.Service

	result := r.db.First(&service, serviceId)
	if result.Error != nil {
		return fmt.Errorf("ServicePostgres - Update - r.db.First: %w", result.Error)
	}

	result = r.db.Model(service).Updates(inputOffice)
	if result.Error != nil {
		return fmt.Errorf("ServicePostgres - Update - r.db.Model().Updates: %w", result.Error)
	}

	return nil
}

func (r *ServicePostgres) Delete(serviceId int64) error {
	result := r.db.Delete(&entity.Service{}, serviceId)
	if result.Error != nil {
		return fmt.Errorf("ServicePostgres - Delete - r.db.Delete: %w", result.Error)
	}

	return nil
}