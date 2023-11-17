package controllers

import (
	"fmt"

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
	fmt.Println(incomeFilter)
	expensesType := ic.IncomeService.GetIncomes(incomeFilter)

	context.JSON(200, gin.H{
		"incomes": expensesType,
	})
}
