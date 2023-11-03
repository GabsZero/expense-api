package repositories

import (
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Repository struct {
}

func (r *Repository) StartDabase() {
	dsn := "host=localhost user=postgres password=senha dbname=expenses-api port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	databaseConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = databaseConnection

	db.AutoMigrate(models.ExpenseType{})
	db.AutoMigrate(models.Expense{})
}
