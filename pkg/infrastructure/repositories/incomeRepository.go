package repositories

import (
	"time"

	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
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

func (ir *IncomeRepository) GetIncomesByMonth(filter dtos.FilterIncomeMonthDto) []models.Income {
	result := []models.Income{}
	if db == nil {
		panic("no db instance found")
	}

	query := db.Model(&result).Preload("IncomeType")
	if !filter.Date.IsZero() {
		startingDate := filter.Date.Format("2006-01-02 15:04:05")
		tempDate := filter.Date.AddDate(0, 1, -filter.Date.Day())
		tempDate = tempDate.Add(time.Hour * 23)
		tempDate = tempDate.Add(time.Minute * 59)
		tempDate = tempDate.Add(time.Second * 59)

		endingDate := tempDate.Format("2006-01-02 15:04:05")

		query.Where("date BETWEEN ? and ?", startingDate, endingDate)

	}

	query.Find(&result)

	return result
}

func (ir *IncomeRepository) Store(newIncome dtos.StoreIncome) models.Income {
	date, err := time.ParseInLocation("2006-01-02", newIncome.Date, time.Local)

	if err != nil {
		panic(err)
	}

	income := models.Income{
		Name:         newIncome.Name,
		Date:         date,
		IncomeTypeId: newIncome.IncomeTypeId,
		Amount:       newIncome.Amount,
	}

	if db == nil {
		panic("no db instance found")
	}

	db.Create(&income)

	return income
}
