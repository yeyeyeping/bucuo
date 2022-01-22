package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/model/request"
	"bucuo/service"
	"bucuo/util/setting"
	"bucuo/util/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
)

var commonservice service.ICommonService = service.CommonService{}

type CommonController struct {
	BaseController
}

func (controller CommonController) Publish(ctx *gin.Context) {
	post := &request.CommonPostReq{}
	if ok := controller.ParseAndValidate(ctx, post); !ok {
		ctx.Abort()
		return
	}
	tbname := ValidateColumn(post.Column, post.Type, false)
	if !tbname {
		controller.CustomerError(ctx, errormsg.ValidateError, "帖子栏目不属于该分类", nil)
		ctx.Abort()
		return
	}
	uid := parseUid(ctx)
	post.PublishId = uint(uid)
	es := commonservice.PushPost(post)
	if es != "" {
		controller.CustomerError(ctx, errormsg.BadRequest, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommonController) FindALl(ctx *gin.Context) {
	param := &request.ByPageCommon{}
	if ok := controller.ParseAndValidate(ctx, param); !ok {
		return
	}
	tbname := ValidateColumn(param.Column, param.Type, true)
	if !tbname {
		controller.CustomerError(ctx, errormsg.ValidateError, "帖子栏目不属于该分类", nil)
		return
	}
	s, rs := commonservice.FindAll(param.Type, param.Column, param.PageSize, param.PageNum)
	if s != "" {
		controller.InternalServerError(ctx, s, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, rs)

}
func (controller CommonController) Detele(ctx *gin.Context) {
	param := &request.DeleteCommonReq{}
	if ok := controller.ParseAndValidate(ctx, param); !ok {
		return
	}
	uid := parseUid(ctx)
	s := commonservice.Delete(param.Type, param.PostID, uint(uid))
	if s != "" {
		controller.CustomerError(ctx, errormsg.UnknowError, s, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller CommonController) FindDetail(ctx *gin.Context) {
	param := &request.DeleteCommonReq{}
	err := ctx.ShouldBindQuery(param)
	if err != nil {
		if err != nil {
			controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
			ctx.Abort()
			return
		}
	}
	s, ok := validator.Validate(param)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return
	}
	res, es := commonservice.FindDetail(param)
	if es != "" {
		controller.InternalServerError(ctx, es, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, res)
}
func ValidateColumn(column string, tablename string, isnull bool) bool {
	if isnull && column == "" {
		return true
	}
	switch tablename {
	case "local_posts":
		if utils.Contains(setting.LocalColumns, column) {
			return true
		}
		break
	case "skill_posts":
		if utils.Contains(setting.SkillColumns, column) {
			return true
		}
		break
	default:
		break
	}
	return false
}
