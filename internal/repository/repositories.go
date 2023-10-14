package repository

import (
	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"gorm.io/gorm"
)

type OfficeRepo interface {
	Create(office entity.Office) (int64, error)
	GetAll() ([]entity.Office, error)
	GetById(officeId int64) (entity.Office, error)
	Update(officeId int64, inputOffice entity.Office) error
	Delete(officeId int64) error
}

type AtmRepo interface {
	Create(atm entity.Atm) (int64, error)
	GetAll() ([]entity.Atm, error)
	GetById(atmId int64) (entity.Atm, error)
	Update(atmId int64, inputAtm entity.Atm) error
	Delete(atmId int64) error
}

type ChartRepo interface {
	Create(chart entity.ChartsOffices) (int64, error)
	GetAll() ([]entity.ChartsOffices, error)
	GetById(chartId int64) (entity.ChartsOffices, error)
	Update(chartId int64, inputChart entity.ChartsOffices) error
	Delete(chartId int64) error
}

type ServiceRepo interface {
	Create(service entity.Service) (int64, error)
	GetAll() ([]entity.Service, error)
	GetById(serviceId int64) (entity.Service, error)
	Update(serviceId int64, inputService entity.Service) error
	Delete(serviceId int64) error
}

type WorkTimeRepo interface {
	Create(workTime entity.WorksTimeOffice) (int64, error)
	GetAll() ([]entity.WorksTimeOffice, error)
	GetById(workTimeId int64) (entity.WorksTimeOffice, error)
	Update(workTimeId int64, inputWorkTime entity.WorksTimeOffice) error
	Delete(workTimeId int64) error
}

type AlgorithmRepo interface {
}

type Repositories struct {
	OfficeRepo
	AtmRepo
	ChartRepo
	ServiceRepo
	WorkTimeRepo
	AlgorithmRepo
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		OfficeRepo: NewOfficePostgres(db),
		AtmRepo: NewAtmPostgres(db),
		ChartRepo: NewChartPostgres(db),
		ServiceRepo: NewServicePostgres(db),
		WorkTimeRepo: NewWorkTimePostgres(db),
		AlgorithmRepo: NewAlgorithmPostgres(db),
	}
}