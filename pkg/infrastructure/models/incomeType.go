package models

import (
	"time"

	"gorm.io/gorm"
)

type IncomeType struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
