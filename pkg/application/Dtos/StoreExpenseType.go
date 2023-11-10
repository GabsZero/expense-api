package dtos

type StoreExpenseTypeDto struct {
	Name string `validate:"required" json: "name" form: "name"`
}
