package dtos

type StoreExpenseTypeDto struct {
	Name string `json:"name" form:"name" validation:"required"`
}
