package controllers

import (
	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
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
func (etc *ExpenseController) Store(context *gin.Context) {
	enableCors(context)
	newExpense := dtos.StoreExpense{}

	err := context.Bind(&newExpense)

	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	expense := etc.ExpenseService.StoreExpense(newExpense)

	context.JSON(200, gin.H{
		"expense": expense,
	})
}

func (etc *ExpenseController) FindExpensesInAMonth(context *gin.Context) {
	enableCors(context)
	monthFilter := dtos.FilterExpenseMonthDto{}

	err := context.Bind(&monthFilter)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	expenses := etc.ExpenseService.FilterExpensesInAMonth(monthFilter)

	context.JSON(200, gin.H{
		"expenses": expenses,
	})
}
