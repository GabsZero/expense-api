package repositories

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
)

type IncomeRepository struct {
}

func (ir *IncomeRepository) GetAll() []models.Income {
	incomes := []models.Income{}
	if db == nil {
		panic("no db instance found")
	}

	db.Model(&incomes).Preload("IncomeType").Find(&incomes)

	return incomes
}
