package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BaseController struct {
}

func (r BaseController) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(errormsg.Success, response.RespModel(0, "", data))
}
func (r BaseController) BadRequest(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(errormsg.BadRequest, response.RespModel(errormsg.BadRequest, msg, data))
}
func (r BaseController) InternalServerError(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(errormsg.InternalServerError, response.RespModel(errormsg.InternalServerError, msg, data))
}
func (r BaseController) CustomerError(ctx *gin.Context, code int, msg string, data interface{}) {
	ctx.JSON(errormsg.BadRequest, response.RespModel(code, msg, data))
}
func (r BaseController) UnkonwError(ctx *gin.Context) {
	ctx.JSON(errormsg.BadRequest, response.RespModel(errormsg.UnknowError, "", nil))
}
func parseUid(ctx *gin.Context) uint64 {
	suserid, _ := ctx.Get("UserId")
	userid, _ := strconv.ParseUint(suserid.(string), 10, 64)
	return userid
}
