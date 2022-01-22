package routes

import (
	"bucuo/controller"
	"github.com/gin-gonic/gin"
)

var usercontroller = controller.UserController{}

func UserRoutesInit(router *gin.RouterGroup) {

	router.POST("user", usercontroller.CreateUser)
	router.GET("user/wxLogin", usercontroller.LoginFromWechat)
	router.POST("user/login", usercontroller.Login)
}
func AuthUserRoutesInit(router *gin.RouterGroup) {
	router.GET("user", usercontroller.GetUser)
	router.PUT("user", usercontroller.UpdateUser)
	router.PUT("user/like/:uid", usercontroller.Like)
	router.DELETE("user/unlike/:uid", usercontroller.UnLike)
	router.GET("user/testLogin", usercontroller.TestLogin)

}
