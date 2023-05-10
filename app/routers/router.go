package routers

import (
	"app/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {

	userGroup := router.Group("/api")
	{
		userGroup.GET("", handlers.GetRoot)
		userGroup.GET("/stock_history", handlers.GetStocks)
	}

	// 404 Not Foundのハンドリング
	router.NoRoute(func(c *gin.Context) {
		c.File(".../client/dist/index.html")
	})

}
