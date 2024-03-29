package routes

import (
	"github.com/gabszero/expenses-api/pkg/application/routes/expenses"
	"github.com/gabszero/expenses-api/pkg/application/routes/expensesType"
	"github.com/gabszero/expenses-api/pkg/application/routes/incomeTypes"
	"github.com/gabszero/expenses-api/pkg/application/routes/incomes"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r *Router) StartRoutes() {
	router := gin.Default()

	expensesRoutes := expenses.ExpensesRoutes{}
	expensesRoutes.DefineRoutes(router)

	incomeRoutes := incomes.IncomeRoutes{}
	incomeRoutes.DefineRoutes(router)

	expensesTypeRoutes := expensesType.ExpensesTypeRoutes{}
	expensesTypeRoutes.DefineRoutes(router)

	incomeTypesRoutes := incomeTypes.IncomeTypesRoutes{}
	incomeTypesRoutes.DefineRoutes(router)

	router.Run() // listen and serve on 0.0.0.0:8080
}
