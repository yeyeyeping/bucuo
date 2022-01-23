package routes

import (
	"bucuo/controller"
	"github.com/gin-gonic/gin"
)

var commoncontroller = controller.CommonController{}

func CommonRoutesInit(router *gin.RouterGroup) {
	router.GET("common/one", commoncontroller.FindDetail)
}

func AuthCommonRoutesInit(router *gin.RouterGroup) {
	router.GET("common", commoncontroller.FindALl)
	router.POST("common", commoncontroller.Publish)
	router.DELETE("common", commoncontroller.Detele)
	router.PUT("common", commoncontroller.Update)
}
