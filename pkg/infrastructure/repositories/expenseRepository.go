package repositories

import (
	"time"

	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
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
		// filter.Date.Format("2006-01-02")
		startingDate := filter.Date.Format("2006-01-02 15:04:05")
		endingDate := filter.Date.Add(24*time.Hour - 1).Format("2006-01-02 15:04:05")

		query.Where("date BETWEEN ? and ?", startingDate, endingDate)
	}

	query.Find(&result)

	return result
}

func (etr *ExpenseRepository) Store(newExpense dtos.StoreExpense) models.Expense {
	date, err := time.ParseInLocation("2006-01-02", newExpense.Date, time.Local)

	if err != nil {
		panic(err)
	}

	expense := models.Expense{
		Name:               newExpense.Name,
		Date:               date,
		CurrentInstallment: newExpense.CurrentInstallment,
		TotalInstallments:  newExpense.TotalInstallments,
		ExpenseTypeId:      newExpense.ExpenseTypeId,
		Amount:             newExpense.Amount,
		IsPaid:             *newExpense.IsPaid,
		IsRecurring:        *newExpense.IsRecurring,
	}

	if db == nil {
		panic("no db instance found")
	}

	db.Create(&expense)

	return expense
}

func (etr *ExpenseRepository) GetExpensesByMonth(filter dtos.FilterExpenseMonthDto) []models.Expense {
	result := []models.Expense{}
	if db == nil {
		panic("no db instance found")
	}

	query := db.Model(&result).Preload("ExpenseType")
	if !filter.Date.IsZero() {
		startingDate := filter.Date.Format("2006-01-02 15:04:05")
		tempDate := filter.Date.AddDate(0, 1, -filter.Date.Day())
		tempDate = tempDate.Add(time.Hour * 23)
		tempDate = tempDate.Add(time.Minute * 59)
		tempDate = tempDate.Add(time.Second * 59)

		endingDate := tempDate.Format("2006-01-02 15:04:05")

		query.Where("date BETWEEN ? and ?", startingDate, endingDate)

	}

	query.Statement.Order("Date").Find(&result)

	return result
}
