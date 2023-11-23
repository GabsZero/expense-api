package controllers

import (
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	ExpenseService services.ExpenseService
}

func (etc *ExpenseController) Index(context *gin.Context) {
	enableCors(context)
	expenseFilter := models.Expense{}
	err := context.Bind(&expenseFilter)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	expenses := etc.ExpenseService.GetExpenses(expenseFilter)

	context.JSON(200, gin.H{
		"expenses": expenses,
	})
}
