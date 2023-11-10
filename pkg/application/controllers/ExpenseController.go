package controllers

import (
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	ExpenseService services.ExpenseService
}

func (etc *ExpenseController) Index(context *gin.Context) {
	expensesType := etc.ExpenseService.GetExpenses()

	context.JSON(200, gin.H{
		"expenses": expensesType,
	})
}
