package repositories

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
)

type IncomeTypesRepository struct {
}

func (itr *IncomeTypesRepository) GetAll() []models.IncomeType {
	incomeTypes := []models.IncomeType{}
	if db == nil {
		panic("no db instance found")
	}

	db.Find(&incomeTypes)

	return incomeTypes
}
