package handlers

import (
	"app/repositories"
	"app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /api/
func GetRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "AAA",
	})
}

func GetStocks(c *gin.Context) {
	stocks, err := repositories.GetStocks(utils.GetDB())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stocks)
}
