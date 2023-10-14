package service

import (
	"github.com/End-rey/VTBMoreTech5Backend/internal/dto"
	"github.com/End-rey/VTBMoreTech5Backend/internal/repository"
)

type Office interface {
	Create(officeDTO dto.OfficeDTO) (int64, error)
	GetAll() ([]dto.OfficeDTO, error)
	GetById(officeId int64) (dto.OfficeDTO, error)
	Delete(officeId int64) error
	Update(officeId int64, input dto.OfficeDTO) error
}

type Atm interface {
	Create(atmDTO dto.AtmDTO) (int64, error)
	GetAll() ([]dto.AtmDTO, error)
	GetById(atmId int64) (dto.AtmDTO, error)
	Delete(atmId int64) error
	Update(atmId int64, input dto.AtmDTO) error
}

type Algorithm interface {
}

type Services struct {
	Office
	Atm
	Algorithm
}

func NewService(repos *repository.Repositories) *Services {
	return &Services{
		Office:    NewOfficeService(*repos),
		Atm:       NewAtmService(repos.AtmRepo, repos.ServiceRepo),
		Algorithm: NewAlgorithmService(repos.AlgorithmRepo),
	}
}
