package services

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

type ExpenseTypeService struct {
	ExpenseTypeRepository *repositories.ExpenseTypeRepository
}

func (ets *ExpenseTypeService) GetExpensesType() []models.ExpenseType {
	return ets.ExpenseTypeRepository.GetAll()
}
