package controllers

import (
	"github.com/gabszero/expenses-api/pkg/domain/services"
	"github.com/gin-gonic/gin"
)

type IncomeTypesController struct {
	IncomeTypeService services.IncomeTypesService
}

func (itc *IncomeTypesController) GetIncomeTypes(context *gin.Context) {
	enableCors(context)
	incomeTypes := itc.IncomeTypeService.GetIncomeTypes()

	context.JSON(200, gin.H{
		"incomeTypes": incomeTypes,
	})
}
