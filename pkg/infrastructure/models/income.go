package models

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `form:"name"`
	Date         time.Time `json:"date" form:"date" time_format:"2006-01-02"`
	IncomeTypeId uint
	Amount       float32 `form:"amount"`
	IncomeType   IncomeType
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
