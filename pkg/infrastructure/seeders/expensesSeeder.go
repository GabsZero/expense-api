package seeders

import (
	"time"

	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

func Populate() {
	populateExpanses()
	populateIncomes()
}

func populateIncomes() {
	date, err := time.Parse("2006-01-02", "2023-11-01")
	if err != nil {
		panic(err)
	}
	incomes := []models.Income{
		{
			Name:         "Entrada de teste 1",
			Date:         date,
			Amount:       5600,
			IncomeTypeId: 1,
		},
	}

	repository := repositories.Repository{}
	db := repository.GetDbInstance()
	for _, income := range incomes {
		exist := models.Income{}
		result := db.Where("name = ?", income.Name).First(&exist)
		if result.RowsAffected == 0 {
			db.Create(&income)
		}
	}
}

func populateExpanses() {
	date, err := time.Parse("2006-01-02", "2023-11-01")
	if err != nil {
		panic(err)
	}
	expenses := []models.Expense{
		{
			Name:               "custo de teste 1",
			Date:               date,
			CurrentInstallment: 1,
			TotalInstallments:  5,
			ExpenseTypeId:      1,
			Amount:             1500.11,
			IsPaid:             false,
			IsRecurring:        false,
		},
		{
			Name:               "Custo de teste 2",
			Date:               date,
			CurrentInstallment: 1,
			TotalInstallments:  5,
			ExpenseTypeId:      1,
			Amount:             1500.11,
			IsPaid:             false,
			IsRecurring:        true,
		},
		{
			Name:               "Custo de teste 3",
			Date:               date,
			CurrentInstallment: 1,
			TotalInstallments:  5,
			ExpenseTypeId:      1,
			Amount:             1500.11,
			IsPaid:             true,
			IsRecurring:        true,
		},
	}

	repository := repositories.Repository{}
	db := repository.GetDbInstance()
	for _, expense := range expenses {
		exist := models.Expense{}
		result := db.Where("name = ?", expense.Name).First(&exist)
		if result.RowsAffected == 0 {
			db.Create(&expense)
		}
	}
}
