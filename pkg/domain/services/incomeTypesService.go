package services

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

type IncomeTypesService struct {
	IncomeTypesRepository *repositories.IncomeTypesRepository
}

func (its *IncomeTypesService) GetIncomeTypes() []models.IncomeType {
	return its.IncomeTypesRepository.GetAll()
}
