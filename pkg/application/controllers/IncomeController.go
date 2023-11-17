package controllers

import (
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gin-gonic/gin"
)

type IncomeController struct {
	IncomeService services.IncomeService
}

func (ic *IncomeController) Index(context *gin.Context) {
	enableCors(context)
	expensesType := ic.IncomeService.GetIncomes()

	context.JSON(200, gin.H{
		"expenses": expensesType,
	})
}
