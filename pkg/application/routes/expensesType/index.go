package expensesType

import (
	"github.com/gabszero/expenses-api/pkg/application/controllers"
	"github.com/gin-gonic/gin"
)

type ExpensesTypeRoutes struct {
}

func (etr *ExpensesTypeRoutes) DefineRoutes(router *gin.Engine) {
	expenseTypeController := controllers.ExpenseTypeController{}
	router.GET("/getExpensesType", expenseTypeController.GetExpensesType)
	router.POST("/expenseType", expenseTypeController.Store)

}
