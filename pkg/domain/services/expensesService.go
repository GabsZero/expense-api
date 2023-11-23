package services

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

type ExpenseService struct {
	ExpenseRepository *repositories.ExpenseRepository
}

func (ets *ExpenseService) GetExpenses(filter models.Expense) []models.Expense {
	return ets.ExpenseRepository.GetAll(filter)
}
