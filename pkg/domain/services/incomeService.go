package services

import (
	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

type IncomeService struct {
	IncomeRepository *repositories.IncomeRepository
}

func (is *IncomeService) GetIncomes(incomeFilter models.Income) []models.Income {
	return is.IncomeRepository.GetAll(incomeFilter)
}

func (is *IncomeService) GetIncomesByMonth(filter dtos.FilterIncomeMonthDto) []models.Income {
	return is.IncomeRepository.GetIncomesByMonth(filter)
}

func (is *IncomeService) StoreIncome(newIncome dtos.StoreIncome) models.Income {
	return is.IncomeRepository.Store(newIncome)
}
