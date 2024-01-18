package incomes

import (
	"github.com/gabszero/expenses-api/pkg/application/controllers"
	"github.com/gin-gonic/gin"
)

type IncomeRoutes struct {
}

func (ir *IncomeRoutes) DefineRoutes(router *gin.Engine) {
	incomeController := controllers.IncomeController{}
	router.GET("/incomes", incomeController.Index)
	router.POST("/incomes", incomeController.Store)
	router.GET("/getIncomesByMonth", incomeController.FindIncomesInAMonth)

}
