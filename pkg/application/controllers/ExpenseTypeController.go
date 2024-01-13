package controllers

import (
	dtos "github.com/gabszero/expenses-api/pkg/application/Dtos"
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ExpenseTypeController struct {
	ExpenseTypeService services.ExpenseTypeService
}

func (etc *ExpenseTypeController) GetExpensesType(context *gin.Context) {
	enableCors(context)
	expensesType := etc.ExpenseTypeService.GetExpensesType()

	context.JSON(200, gin.H{
		"expensesType": expensesType,
	})
}

func (etc *ExpenseTypeController) Store(context *gin.Context) {
	// validar a requisição
	storeExpenseTypeDto := dtos.StoreExpenseTypeDto{}
	context.Bind(&storeExpenseTypeDto)
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.Struct(storeExpenseTypeDto)

	context.JSON(201, gin.H{
		"expensesType": storeExpenseTypeDto,
	})
}
