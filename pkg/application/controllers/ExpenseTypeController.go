package controllers

import (
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gin-gonic/gin"
)

type ExpenseTypeController struct {
	ExpenseTypeService services.ExpenseTypeService
}

func (etc *ExpenseTypeController) GetExpensesType(context *gin.Context) {
	expensesType := etc.ExpenseTypeService.GetExpensesType()

	context.JSON(200, gin.H{
		"expensesTYpe": expensesType,
	})
}
