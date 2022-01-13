package controller

import (
	"bucuo/middleware"
	"bucuo/model"
	"bucuo/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	BaseController
}

func (uc UserController) Login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	if ok := (service.UserService{}.FindByUserName(&user)); ok {
		log.Printf("%#v", user)
		if jwtstring, err := middleware.GenerateJwt(user.Id); err != nil {
			uc.BadRequest(ctx, err.Error(), nil)
		} else {
			uc.Success(ctx, gin.H{
				"token": jwtstring,
			})
		}
	} else {
		uc.BadRequest(ctx, "Password error", nil)
	}

}
