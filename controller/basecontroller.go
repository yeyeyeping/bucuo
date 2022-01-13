package controller

import (
	"bucuo/model"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (r *BaseController) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, model.BaseResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}
func (r *BaseController) BadRequest(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(400, model.BaseResponse{
		Code: -1,
		Msg:  msg,
		Data: data,
	})
}
