package services

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

type IncomeService struct {
	IncomeRepository *repositories.IncomeRepository
}

func (is *IncomeService) GetIncomes(incomeFilter models.Income) []models.Income {
	return is.IncomeRepository.GetAll(incomeFilter)
}
