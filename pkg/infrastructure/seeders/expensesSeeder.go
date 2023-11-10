package seeders

import (
	"time"

	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gabszero/expenses-api/pkg/infrastructure/repositories"
)

func Populate() {
	date, err := time.Parse("Jan 02, 2006", "Nov 01, 2023")
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
