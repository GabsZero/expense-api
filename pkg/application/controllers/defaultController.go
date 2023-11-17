package controllers

import "github.com/gin-gonic/gin"

func enableCors(context *gin.Context) {
	context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
}
