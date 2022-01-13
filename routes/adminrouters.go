package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdminRoutesInit(router *gin.RouterGroup) {

	router.GET("admin/ping", func(c *gin.Context) {
		value, _ := c.Get("UserId")
		c.JSON(http.StatusOK, gin.H{
			"ID": value,
		})
	})
}
