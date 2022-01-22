package routes

import (
	"bucuo/controller"
	"github.com/gin-gonic/gin"
)

var exprcontroller controller.ExprController

func AuthExprRoutesInit(router *gin.RouterGroup) {
	router.POST("/exprpost", exprcontroller.Publish)
	router.PUT("/exprpost", exprcontroller.UpdateOne)
	router.DELETE("/exprpost/:id", exprcontroller.DeleteOne)
	router.GET("/exprpost/all", exprcontroller.FindALl)
}
func ExprRoutesInit(router *gin.RouterGroup) {
	router.GET("/exprpost/:id", exprcontroller.FindOne)
}
