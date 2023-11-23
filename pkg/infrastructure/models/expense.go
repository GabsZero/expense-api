package models

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	ID                 uint      `gorm:"primaryKey"`
	Name               string    `form:"name"`
	Date               time.Time `json:"date" form:"date" time_format:"2006-01-02"`
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
