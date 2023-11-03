package repositories

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
)

type ExpenseTypeRepository struct {
}

func (etr *ExpenseTypeRepository) GetAll() []models.ExpenseType {
	expensesType := []models.ExpenseType{}
	if db == nil {
		panic("no db instance found")
	}

	db.Model(&expensesType).Find(&expensesType)

	return expensesType
}
