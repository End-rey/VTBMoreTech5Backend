package repository

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"gorm.io/gorm"
)

type WorkTimePostgres struct {
	db *gorm.DB
}

func NewWorkTimePostgres(db *gorm.DB) *WorkTimePostgres {
	return &WorkTimePostgres{db: db}
}

func (r *WorkTimePostgres) Create(workTime entity.WorksTime) (int64, error) {
	result := r.db.Create(&workTime)
	if result.Error != nil {
		return 0, fmt.Errorf("WorkTimePostgres - Create - r.db.Create: %w", result.Error)
	}

	return workTime.ID, nil
}

func (r *WorkTimePostgres) GetAll() ([]entity.WorksTime, error) {
	var workTimes []entity.WorksTime

	result := r.db.Find(&workTimes)
	if result.Error != nil {
		return nil, fmt.Errorf("WorkTimePostgres - GetAll - r.db.Find: %w", result.Error)
	}

	return workTimes, nil
}

func (r *WorkTimePostgres) GetById(workTimeId int64) (entity.WorksTime, error) {
	var workTime entity.WorksTime

	result := r.db.First(&workTime, workTimeId)
	if result.Error != nil {
		return workTime, fmt.Errorf("WorkTimePostgres - GetById - r.db.First: %w", result.Error)
	}

	return workTime, nil
}

func (r *WorkTimePostgres) Update(workTimeId int64, inputworkTime entity.WorksTime) error {
	var workTime entity.WorksTime

	result := r.db.First(&workTime, workTimeId)
	if result.Error != nil {
		return fmt.Errorf("WorkTimePostgres - Update - r.db.First: %w", result.Error)
	}

	result = r.db.Model(workTime).Updates(inputworkTime)
	if result.Error != nil {
		return fmt.Errorf("WorkTimePostgres - Update - r.db.Model().Updates: %w", result.Error)
	}

	return nil
}

func (r *WorkTimePostgres) Delete(workTimeId int64) error {
	result := r.db.Delete(&entity.WorksTime{}, workTimeId)
	if result.Error != nil {
		return fmt.Errorf("WorkTimePostgres - Delete - r.db.Delete: %w", result.Error)
	}

	return nil
}