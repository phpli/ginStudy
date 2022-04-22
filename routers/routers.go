package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello router",
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("route", helloHandler)
	return router
}
