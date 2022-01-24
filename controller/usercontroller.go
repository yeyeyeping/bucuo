package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/middleware"
	"bucuo/model/request"
	"bucuo/model/response"
	"fmt"

	"bucuo/service"
	"bucuo/util/validator"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	BaseController
}

var userservice service.IUserServiece = service.UserService{}

func (controller UserController) Login(ctx *gin.Context) {
	user := request.UserReq{}
	ctx.ShouldBind(&user)
	s, ok := validator.Validate(user)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return
	}
	err, suc := userservice.LoginByUserName(user.UserName, user.Password)
	if suc {
		//签发jwt
		jwts, jwterr := middleware.GenerateJwt(err)
		if jwterr != nil {
			controller.CustomerError(ctx, errormsg.JwtError, jwterr.Error(), nil)
			ctx.Abort()
		} else {
			controller.Success(ctx, gin.H{
				"token": jwts,
			})
		}
	} else {
		controller.BadRequest(ctx, err, nil)
	}
}
func (controller UserController) TestLogin(c *gin.Context) {
	userid, ok := c.Get("UserId")
	if ok {
		c.JSON(errormsg.Success, response.RespModel(errormsg.Ok, "", gin.H{
			"UserId": userid,
		}))
	} else {
		c.JSON(errormsg.BadRequest, response.RespModel(errormsg.JwtError, "", nil))
	}
}
func (controller UserController) LoginFromWechat(ctx *gin.Context) {
	opeid := ctx.GetHeader("X-WX-OPENID")
	if opeid == "" {
		controller.BadRequest(ctx, "请从微信端登录", nil)
		ctx.Abort()
		return
	}
	ok, r := userservice.ValidateOpenId(opeid)
	if !ok {
		ctx.AbortWithStatusJSON(
			errormsg.BadRequest,
			response.RespModel(
				errormsg.ImperfectPersonalInformation,
				errormsg.GetMsg(errormsg.ImperfectPersonalInformation, ""),
				nil))
		return
	} else {
		jwts, jwterr := middleware.GenerateJwt(r)
		if jwterr != nil {
			controller.CustomerError(ctx, errormsg.JwtError, jwterr.Error(), nil)
			ctx.Abort()
			return
		} else {
			controller.Success(ctx, gin.H{
				"token": jwts,
			})
		}
	}
}
func (controller UserController) CreateUser(ctx *gin.Context) {
	user := request.UserCreateReq{}
	ctx.ShouldBind(&user)
	fmt.Println(ctx.GetRawData())
	s, ok := validator.Validate(user)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return
	}
	opeid := ctx.GetHeader("X-WX-OPENID")
	if opeid == "" {
		controller.BadRequest(ctx, "请从微信端登录", nil)
		ctx.Abort()
		return
	}
	ok, r := userservice.ValidateOpenId(opeid)
	if ok {
		controller.UnkonwError(ctx)
		ctx.Abort()
		return
	}
	user.OpenId = opeid
	ok, s = userservice.CreateUser(&user)
	if !ok {
		controller.BadRequest(ctx, s, nil)
		ctx.Abort()
		return
	}
	jwts, jwterr := middleware.GenerateJwt(r)
	if jwterr != nil {
		controller.CustomerError(ctx, errormsg.JwtError, jwterr.Error(), nil)
		ctx.Abort()
	} else {
		controller.Success(ctx, gin.H{
			"token": jwts,
		})
	}
}
func (controller UserController) GetUser(ctx *gin.Context) {
	uid := parseUid(ctx)
	if uid == 0 {
		controller.UnkonwError(ctx)
		ctx.Abort()
		return
	}
	user := userservice.GetUserById(uint(uid))
	controller.Success(ctx, gin.H{
		"userinfo": user,
	})
}
func (controller UserController) UpdateUser(ctx *gin.Context) {
	user := &request.UserCreateReq{}
	err := ctx.ShouldBind(user)
	s, ok := validator.Validate(user)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return
	}
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	userid := parseUid(ctx)
	err = userservice.UpdateUser(uint(userid), user)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller UserController) Like(ctx *gin.Context) {
	var uid string
	uid = ctx.Param("uid")
	liked, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	userid := parseUid(ctx)
	errs := userservice.AddHomeLike(uint(userid), uint(liked))
	if errs != "" {
		controller.CustomerError(ctx, errormsg.BadRequest, errs, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller UserController) UnLike(ctx *gin.Context) {
	var uid string
	uid = ctx.Param("uid")
	liked, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	userid := parseUid(ctx)
	errs := userservice.DeleteHomeLike(uint(userid), uint(liked))
	if errs != "" {
		controller.CustomerError(ctx, errormsg.BadRequest, errs, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, nil)
}
func (controller UserController) GetOther(ctx *gin.Context) {
	sid := ctx.Param("id")
	id, err := strconv.ParseUint(sid, 10, 64)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, "id错误", nil)
		ctx.Abort()
		return
	}
	uid := uint(parseUid(ctx))
	rs := userservice.GetOtherUserById(uid, uint(id))
	if rs == "" {
		controller.InternalServerError(ctx, "用户不存在", nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, rs)
}
func (controller UserController) GetUserOne(ctx *gin.Context) {
	sid := ctx.Query("id")
	var uid uint
	if sid == "" {
		uid = uint(parseUid(ctx))
	} else {
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			controller.CustomerError(ctx, errormsg.ValidateError, "id错误", nil)
			ctx.Abort()
			return
		}
		uid = uint(id)
	}
	expr, err := userservice.GetUserExpr(uid)
	if err != nil {
		controller.InternalServerError(ctx, err.Error(), nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, expr)
}
func (controller UserController) GetDetail(ctx *gin.Context) {
	sid := ctx.Query("id")
	var uid uint
	if sid == "" {
		uid = uint(parseUid(ctx))
	} else {
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			controller.CustomerError(ctx, errormsg.ValidateError, "id错误", nil)
			ctx.Abort()
			return
		}
		uid = uint(id)
	}
	r := userservice.GetDetail(uid)
	controller.Success(ctx, r)
}
func (controller UserController) FindUser(ctx *gin.Context) {
	sid := ctx.Query("id")
	var uid uint
	if sid == "" {
		uid = uint(parseUid(ctx))
	} else {
		id, err := strconv.ParseUint(sid, 10, 64)
		if err != nil {
			controller.CustomerError(ctx, errormsg.ValidateError, "id错误", nil)
			ctx.Abort()
			return
		}
		uid = uint(id)
	}
	rs, s := commonservice.FindUserSkillPost(uid)
	if s != "" {
		controller.InternalServerError(ctx, s, nil)
		ctx.Abort()
		return
	}
	controller.Success(ctx, rs)
}
