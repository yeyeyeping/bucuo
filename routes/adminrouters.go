package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutesInit(router *gin.Engine) {
	admin := router.Group("/api/admin")
	{
		admin.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}
}
