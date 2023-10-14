package service

import (
	"fmt"

	"github.com/End-rey/VTBMoreTech5Backend/internal/dto"
	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
	"github.com/End-rey/VTBMoreTech5Backend/internal/repository"
	"github.com/End-rey/VTBMoreTech5Backend/internal/utils"
)

type AtmService struct {
	atmRepo repository.AtmRepo
	serviceRepo repository.ServiceRepo
}

func NewAtmService(
	atmRepo repository.AtmRepo,
	serviceRepo repository.ServiceRepo,
) *AtmService {
	return &AtmService{
		atmRepo: atmRepo,
		serviceRepo: serviceRepo,
	}

}

func (os *AtmService) Create(atmDTO dto.AtmDTO) (int64, error) {
	return 0, nil
}

func (os *AtmService) GetAll() ([]dto.AtmDTO, error) {
	var (
		err        error
		atmsDTO []dto.AtmDTO
		atms    []entity.Atm
	)

	atms, err = os.atmRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("AtmService - GetALL - %w", err)
	}

	for _, atm := range atms {
		atmsDTO = append(atmsDTO, utils.AtmToDTO(atm))
	}

	return atmsDTO, nil
}

func (os *AtmService) GetById(atmId int64) (dto.AtmDTO, error) {
	return dto.AtmDTO{}, nil
}

func (os *AtmService) Delete(atmId int64) error {
	return nil
}

func (os *AtmService) Update(atmId int64, input dto.AtmDTO) error {
	return nil
}
