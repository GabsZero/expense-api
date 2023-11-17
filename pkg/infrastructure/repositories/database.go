package repositories

import (
	"fmt"
	"os"

	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Repository struct {
}

func (r *Repository) GetDbInstance() *gorm.DB {
	return db
}

func (r *Repository) StartDabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	host := os.Getenv("EXPENSE_API_DATABASE_HOST")
	user := os.Getenv("EXPENSE_API_DATABASE_USER")
	password := os.Getenv("EXPENSE_API_DATABASE_PASSWORD")
	database := os.Getenv("EXPENSE_API_DATABASE_DATABASE")
	port := os.Getenv("EXPENSE_API_DATABASE_PORT")
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Sao_Paulo", host, user, password, database, port)

	databaseConnection, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = databaseConnection

	db.AutoMigrate(models.ExpenseType{})
	db.AutoMigrate(models.Expense{})
	db.AutoMigrate(models.IncomeType{})
	db.AutoMigrate(models.Income{})
}
