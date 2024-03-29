package dtos

type StoreIncome struct {
	Name         string  `form:"name"  binding:"required"`
	Date         string  `form:"date" binding:"required"`
	IncomeTypeId uint    `form:"incomeTypeId"  binding:"required"`
	Amount       float32 `form:"amount"  binding:"required"`
}
