package models

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Date         time.Time
	IncomeTypeId uint
	Amount       float32
	IncomeType   IncomeType
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
