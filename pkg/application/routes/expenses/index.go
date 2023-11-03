package expenses

import "github.com/gin-gonic/gin"

type ExpensesRoutes struct {
}

func (er *ExpensesRoutes) DefineRoutes(router *gin.Engine) {
	router.GET("/getExpenses", er.getExpenses)
}

func (er *ExpensesRoutes) getExpenses(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong again haha",
	})
}
