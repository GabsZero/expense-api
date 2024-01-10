package dtos

type StoreExpense struct {
	Name               string  `form:"name"  binding:"required"`
	Date               string  `form:"date" binding:"required"`
	CurrentInstallment uint    `form:"currentInstallment" binding:"required"`
	TotalInstallments  uint    `form:"totalInstallments"  binding:"required"`
	ExpenseTypeId      uint    `form:"expenseTypeId"  binding:"required"`
	Amount             float32 `form:"amount"  binding:"required"`
	IsPaid             *bool   `form:"isPaid"  binding:"required"`
	IsRecurring        *bool   `form:"isRecurring"  binding:"required"`
}
