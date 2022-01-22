package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/model/response"
	"bucuo/util/validator"
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
func (controller BaseController) ParseAndValidate(ctx *gin.Context, data interface{}) bool {
	err := ctx.ShouldBind(data)
	if err != nil {
		if err != nil {
			controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
			ctx.Abort()
			return false
		}
	}
	s, ok := validator.Validate(data)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return false
	}
	return true
}
