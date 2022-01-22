package routes

import (
	"bucuo/controller"
	"github.com/gin-gonic/gin"
)

var rc = controller.ResourceController{}

func AuthResouceRoutesInit(group *gin.RouterGroup) {
	group.POST("resource", rc.Upload)
}
func ResouceRoutesInit(group *gin.RouterGroup) {
	group.GET("resource/:uuid", rc.GetResouce)
}
