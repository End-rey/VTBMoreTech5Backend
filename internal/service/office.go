package service

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/dto"
	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"github.com/End-rey/VTBMoreTech5Backend/internal/repository"
	"github.com/End-rey/VTBMoreTech5Backend/internal/utils"
)

type OfficeService struct {
	repos repository.Repositories
}

func NewOfficeService(
	repos repository.Repositories,
) (*OfficeService) {
	return &OfficeService{
		repos: repos,
	}

}

func (os *OfficeService) Create(officeDTO dto.OfficeDTO) (int64, error) {
	return 0, nil
}

func (os *OfficeService) GetAll() ([]dto.OfficeDTO, error) {
	var (
		err error;
		officesDTO []dto.OfficeDTO;
		offices []entity.Office;
	)

	offices, err = os.repos.OfficeRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("OfficeService - GetALL - %w", err)
	}

	for _, office := range offices {
		officesDTO = append(officesDTO, utils.OfficeToDTO(office))
	}

	return officesDTO, nil
}

func (os *OfficeService) GetById(officeId int64) (dto.OfficeDTO, error) {
	return dto.OfficeDTO{}, nil
}

func (os *OfficeService) Delete(officeId int64) error {
	return nil
}

func (os *OfficeService) Update(officeId int64, input dto.OfficeDTO) error {
	return nil
}

