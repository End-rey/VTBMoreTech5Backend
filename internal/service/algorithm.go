package service

import "github.com/End-rey/VTBMoreTech5Backend/internal/repository"

type AlgorithmRepo struct {
	repo repository.AlgorithmRepo
}

func NewAlgorithmService(repo repository.AlgorithmRepo) *AlgorithmRepo {
	return &AlgorithmRepo{repo: repo}
}