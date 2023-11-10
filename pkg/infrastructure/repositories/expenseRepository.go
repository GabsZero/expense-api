package repositories

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
)

type ExpenseRepository struct {
}

func (etr *ExpenseRepository) GetAll() []models.Expense {
	expenses := []models.Expense{}
	if db == nil {
		panic("no db instance found")
	}

	db.Model(&expenses).Find(&expenses)

	return expenses
}
