package routes

import (
	"bucuo/controller"

	"github.com/gin-gonic/gin"
)

func UserRoutesInit(router *gin.RouterGroup) {
	router.POST("user/resister", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "register success",
		})
	})
	router.GET("user/login", controller.UserController{}.Login)
}
func AuthUserRoutesInit(router *gin.RouterGroup) {

}
