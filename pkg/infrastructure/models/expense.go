package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID                 uint `gorm:"primaryKey"`
	Name               string
	Date               time.Time
	CurrentInstallment uint
	TotalInstallments  uint
	ExpenseTypeId      uint
	Amount             float32
	ExpenseType        ExpenseType
	IsPaid             bool
	IsRecurring        bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
