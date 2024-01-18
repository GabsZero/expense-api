package controllers

import (
	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gabszero/expenses-api/pkg/infrastructure/models"
	"github.com/gin-gonic/gin"
)

type IncomeController struct {
	IncomeService services.IncomeService
}

func (ic *IncomeController) Index(context *gin.Context) {
	enableCors(context)
	incomeFilter := models.Income{}
	err := context.Bind(&incomeFilter)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	incomes := ic.IncomeService.GetIncomes(incomeFilter)

	context.JSON(200, gin.H{
		"incomes": incomes,
	})
}

func (ic *IncomeController) Store(context *gin.Context) {
	enableCors(context)
	newIncome := dtos.StoreIncome{}

	err := context.Bind(&newIncome)

	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	expense := ic.IncomeService.StoreIncome(newIncome)

	context.JSON(200, gin.H{
		"expense": expense,
	})
}

func (ic *IncomeController) FindIncomesInAMonth(context *gin.Context) {
	enableCors(context)
	monthFilter := dtos.FilterIncomeMonthDto{}

	err := context.Bind(&monthFilter)
	if err != nil {
		context.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	incomes := ic.IncomeService.GetIncomesByMonth(monthFilter)

	context.JSON(200, gin.H{
		"incomes": incomes,
	})
}
