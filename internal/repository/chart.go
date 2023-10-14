package repository

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"gorm.io/gorm"
)

type ChartPostgres struct {
	db *gorm.DB
}

func NewChartPostgres(db *gorm.DB) *ChartPostgres {
	return &ChartPostgres{db: db}
}

func (r *ChartPostgres) Create(chart entity.Chart) (int64, error) {
	result := r.db.Create(&chart)
	if result.Error != nil {
		return 0, fmt.Errorf("ChartPostgres - Create - r.db.Create: %w", result.Error)
	}

	return chart.ID, nil
}

func (r *ChartPostgres) GetAll() ([]entity.Chart, error) {
	var charts []entity.Chart

	result := r.db.Find(&charts)
	if result.Error != nil {
		return nil, fmt.Errorf("ChartPostgres - GetAll - r.db.Find: %w", result.Error)
	}

	return charts, nil
}

func (r *ChartPostgres) GetById(chartId int64) (entity.Chart, error) {
	var chart entity.Chart

	result := r.db.First(&chart, chartId)
	if result.Error != nil {
		return chart, fmt.Errorf("ChartPostgres - GetById - r.db.First: %w", result.Error)
	}

	return chart, nil
}

func (r *ChartPostgres) Update(chartId int64, inputchart entity.Chart) error {
	var chart entity.Chart

	result := r.db.First(&chart, chartId)
	if result.Error != nil {
		return fmt.Errorf("ChartPostgres - Update - r.db.First: %w", result.Error)
	}

	result = r.db.Model(chart).Updates(inputchart)
	if result.Error != nil {
		return fmt.Errorf("ChartPostgres - Update - r.db.Model().Updates: %w", result.Error)
	}

	return nil
}

func (r *ChartPostgres) Delete(chartId int64) error {
	result := r.db.Delete(&entity.Chart{}, chartId)
	if result.Error != nil {
		return fmt.Errorf("ChartPostgres - Delete - r.db.Delete: %w", result.Error)
	}

	return nil
}