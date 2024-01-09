package services

import (
	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

type ExpenseService struct {
	ExpenseRepository *repositories.ExpenseRepository
}

func (ets *ExpenseService) GetExpenses(filter models.Expense) []models.Expense {
	return ets.ExpenseRepository.GetAll(filter)
}

func (ets *ExpenseService) FilterExpensesInAMonth(filter dtos.FilterExpenseMonthDto) []models.Expense {
	return ets.ExpenseRepository.GetExpensesByMonth(filter)
}
