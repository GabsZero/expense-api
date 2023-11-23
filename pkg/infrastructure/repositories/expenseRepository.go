package repositories

import (
	"time"

	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
)

type ExpenseRepository struct {
}

func (etr *ExpenseRepository) GetAll(filter models.Expense) []models.Expense {
	result := []models.Expense{}
	if db == nil {
		panic("no db instance found")
	}

	query := db.Model(&result).Preload("ExpenseType")
	if !filter.Date.IsZero() {
		dateParsed, err := time.Parse("2006-01-02", filter.Date.Format("2006-01-02"))
		if err != nil {
			panic(err)
		}

		query.Where("date = ?", dateParsed)
	}

	query.Find(&result)

	return result
}
