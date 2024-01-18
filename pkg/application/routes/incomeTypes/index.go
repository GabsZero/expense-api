package incomeTypes

import (
	"github.com/gabszero/expenses-api/pkg/application/controllers"
	"github.com/gin-gonic/gin"
)

type IncomeTypesRoutes struct {
}

func (itr *IncomeTypesRoutes) DefineRoutes(router *gin.Engine) {
	incomeTypesController := controllers.IncomeTypesController{}
	router.GET("/getIncomeTypes", incomeTypesController.GetIncomeTypes)

}
