package routes

import (
	"bucuo/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutesInit(router *gin.Engine) {
	user := router.Group("/api/user")
	{
		user.POST("/resister", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "register success",
			})
		})
		user.GET("/login", controller.UserController{}.Login)
	}
}
