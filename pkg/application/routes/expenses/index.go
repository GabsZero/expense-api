package expenses

import (
	"github.com/gabszero/expenses-api/pkg/application/controllers"
	"github.com/gin-gonic/gin"
)

type ExpensesRoutes struct {
}

func (er *ExpensesRoutes) DefineRoutes(router *gin.Engine) {
	expenseController := controllers.ExpenseController{}
	router.GET("/expenses", expenseController.Index)
	router.GET("/getExpensesByMonth", expenseController.FindExpensesInAMonth)
}
