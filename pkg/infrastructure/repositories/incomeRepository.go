package repositories

import (
	"time"

	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
)

type IncomeRepository struct {
}

func (ir *IncomeRepository) GetAll(incomeFilter models.Income) []models.Income {
	result := []models.Income{}
	if db == nil {
		panic("no db instance found")
	}

	query := db.Model(&result).Preload("IncomeType")
	if !incomeFilter.Date.IsZero() {
		dateParsed, err := time.Parse("2006-01-02", incomeFilter.Date.Format("2006-01-02"))
		if err != nil {
			panic(err)
		}

		query.Where("date = ?", dateParsed)
	}

	query.Find(&result)

	return result
}
