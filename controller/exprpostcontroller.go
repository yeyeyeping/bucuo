package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/model/request"
	"bucuo/service"
	"bucuo/util/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ExprController struct {
	BaseController
}

var exprservice service.IExprPostService = service.ExprPostService{}

func (controller ExprController) Publish(ctx *gin.Context) {
	post := &request.ExprPostReq{}
	err := ctx.ShouldBind(post)
	if err != nil {
		if err != nil {
			controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
			ctx.Abort()
			return
		}
	}
	s, ok := validator.Validate(post)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return
	}
	uid := parseUid(ctx)
	post.PublishId = uint(uid)
	es := exprservice.PushPost(post)
	if es != "" {
		if err != nil {
			controller.CustomerError(ctx, errormsg.BadRequest, es, nil)
			ctx.Abort()
			return
		}
	}
	controller.Success(ctx, nil)
}
func (controller ExprController) FindALl(ctx *gin.Context) {
	pageparam := &request.ByPage{}
	if ok := controller.ParseAndValidate(ctx, pageparam); !ok {
		return
	}
	result, err := exprservice.FindAll(pageparam.Column, pageparam.PageSize, pageparam.PageNum)
	if err != nil {
		controller.CustomerError(ctx, errormsg.BadRequest, err.Error(), nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, result)
}
func (controller ExprController) FindOne(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	err2, result := exprservice.FindDetails(uint(id))
	if err2 != nil {
		controller.InternalServerError(ctx, err2.Error(), nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, result)
}
func (controller ExprController) DeleteOne(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	uid := parseUid(ctx)
	err = exprservice.ExprExist(uint(id), uint(uid))
	if err != nil {
		controller.CustomerError(ctx, errormsg.InternalServerError, err.Error(), nil)
		ctx.Abort()
		return
	}
	if err = exprservice.DeleteOne(uint(id)); err != nil {
		controller.CustomerError(ctx, errormsg.InternalServerError, err.Error(), nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller ExprController) UpdateOne(ctx *gin.Context) {
	post := &request.UpdateExprReq{}
	if ok := controller.ParseAndValidate(ctx, post); !ok {
		return
	}
	id := parseUid(ctx)
	err := exprservice.ExprExist(post.ID, uint(id))
	if err != nil {
		controller.CustomerError(ctx, errormsg.InternalServerError, err.Error(), nil)
		ctx.Abort()
		return
	}
	err = exprservice.UpdateOne(post)
	if err != nil {
		controller.CustomerError(ctx, errormsg.InternalServerError, err.Error(), nil)
		ctx.Abort()
		return
	}
	if err = exprservice.DeleteOne(uint(id)); err != nil {
		controller.CustomerError(ctx, errormsg.InternalServerError, err.Error(), nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
