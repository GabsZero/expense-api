package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID                 uint `gorm:"primaryKey"`
	Name               string
	CurrentInstallment uint
	TotalInstallments  uint
	ExpenseTypeId      uint
	ExpenseType        ExpenseType
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
