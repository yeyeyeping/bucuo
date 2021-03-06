package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/model/request"
	"bucuo/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CommentController struct {
	BaseController
}

var commentservice service.ICommentService = service.CommentService{}

func (controller CommentController) AddComment(ctx *gin.Context) {
	param := &request.CommentReq{}
	if ok := controller.ParseAndValidate(ctx, param); !ok {
		return
	}
	uid := uint(parseUid(ctx))
	if es := commentservice.AddComment(param, uid); es != "" {
		controller.BadRequest(ctx, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommentController) DeleteComment(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, "id错误", nil)
		ctx.Abort()
		return
	}
	uid := uint(parseUid(ctx))
	es := commentservice.DeleteCommemnt(uint(id), uid)
	if es != "" {
		controller.InternalServerError(ctx, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommentController) AddReply(ctx *gin.Context) {
	s := &request.AddReplyReq{}
	if ok := controller.ParseAndValidate(ctx, s); !ok {
		return
	}
	uid := uint(parseUid(ctx))
	if s := commentservice.AddReply(s, uid); s != "" {
		controller.InternalServerError(ctx, s, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommentController) DeleteReply(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, "id错误", nil)
		ctx.Abort()
		return
	}
	uid := uint(parseUid(ctx))
	es := commentservice.DeleteReply(uint(id), uid)
	if es != "" {
		controller.InternalServerError(ctx, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommentController) Like(ctx *gin.Context) {
	sid := ctx.Param("id")
	uint64id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	id := uint(uint64id)
	uid := parseUid(ctx)
	if es := commentservice.Like(uint(uid), id); es != "" {
		controller.CustomerError(ctx, errormsg.UnknowError, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommentController) LikeReply(ctx *gin.Context) {
	sid := ctx.Param("id")
	uint64id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	id := uint(uint64id)
	uid := parseUid(ctx)
	if es := commentservice.LikeReply(uint(uid), id); es != "" {
		controller.CustomerError(ctx, errormsg.InternalServerError, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommentController) GetCommentByPage(ctx *gin.Context) {
	c := &request.ByPageComment{}
	if ok := controller.ParseAndValidate(ctx, c); !ok {
		return
	}
	es, rs := commentservice.GetComments(c)
	if es != "" {
		controller.InternalServerError(ctx, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, rs)
}
func (controller CommentController) GetReply(ctx *gin.Context) {
	p := &request.ByPageReply{}
	if ok := controller.ParseAndValidate(ctx, p); !ok {
		return
	}
	es, re := commentservice.GetReplies(p)
	if es != "" {
		controller.InternalServerError(ctx, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, re)

}
