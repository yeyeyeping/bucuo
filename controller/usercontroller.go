package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (uc UserController) Login(ctx *gin.Context) {

	uc.Success(ctx, "login success")
}
