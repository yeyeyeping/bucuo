package controller

import (
	"bucuo/constant/errormsg"
	"bucuo/model/request"
	"bucuo/service"
	"bucuo/util/setting"
	"bucuo/util/validator"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/utils"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type ResourceController struct {
	BaseController
}

var iresourceservice service.IResourceService = service.ResourceService{}

func (controller ResourceController) Upload(ctx *gin.Context) {
	rparam := &request.ResourceReq{}
	//绑定参数
	err := ctx.ShouldBindQuery(rparam)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	//验证参数
	s, ok := validator.Validate(rparam)
	if !ok {
		controller.CustomerError(ctx, errormsg.ValidateError, s, nil)
		ctx.Abort()
		return
	}
	//获取uid
	rparam.UserID = uint(parseUid(ctx))
	i1, e := iresourceservice.PostExist(rparam)
	if !i1 {
		controller.CustomerError(ctx, errormsg.ResourceError, "", nil)
		ctx.Abort()
		return
	}

	if e != nil {
		controller.CustomerError(ctx, errormsg.InternalServerError, "", nil)
		ctx.Abort()
		return
	}
	//获取表单
	form, err := ctx.MultipartForm()
	if err != nil {
		controller.BadRequest(ctx, err.Error(), nil)
		ctx.Abort()
		return
	}
	//检查帖子是否合法
	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	var filesid = make([]string, len(files))
	// 检查帖子资源数是否超出
	if i, e := iresourceservice.GetPostCount(rparam); e != nil || int(i)+len(files) > 10 {
		var emsg int
		var es string
		if e != nil {
			emsg = errormsg.InternalServerError
			es = ""
		}
		if int(i)+len(files) > 10 {
			emsg = errormsg.ResourceError
			es = "单个帖子最多上传10个资源！"
		}
		controller.CustomerError(ctx, emsg, es, nil)
		ctx.Abort()
		return
	}
	for ind, file := range files {
		//判断上传资源的大小
		if file.Size > 1024*1024*10 {
			controller.BadRequest(ctx, "上传文件必须小于10MB", nil)
			ctx.Abort()
			return
		}
		//判断文件类型是否合法
		ext := path.Ext(file.Filename)
		if !utils.Contains(setting.Extension, ext) {
			controller.BadRequest(ctx, "文件格式不允许", nil)
			ctx.Abort()
			return
		}
		//生成主键
		u, _ := uuid.NewUUID()
		uid := u.String()
		//添加文件的uuid到列表中
		filesid[ind] = uid
		//生成文件的磁盘路径
		abpath := setting.ResourcePath + string(os.PathSeparator) + uid + ext
		if iresourceservice.CreateResource(uid, abpath, rparam.UserID, rparam.OwnerType, rparam.OwnerID) != nil {
			controller.InternalServerError(ctx, "", nil)
			ctx.Abort()
			return
		}

		if err := ctx.SaveUploadedFile(file, abpath); err != nil {
			ctx.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	controller.Success(ctx, gin.H{
		"resourcelist": filesid,
	})
}
func (controller ResourceController) GetResouce(ctx *gin.Context) {
	suuid := ctx.Param("uuid")
	_, err := uuid.Parse(suuid)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ValidateError, err.Error(), nil)
		ctx.Abort()
		return
	}
	s, err := iresourceservice.GetResouce(suuid)
	if err != nil {
		controller.CustomerError(ctx, errormsg.ResourceError, err.Error(), nil)
		ctx.Abort()
		return
	}
	bs, err := ioutil.ReadFile(s)
	if err != nil {
		controller.InternalServerError(ctx, "", nil)
		ctx.Abort()
		return
	}
	ctx.Writer.Write(bs)
	ctx.Writer.Flush()
}
func (controller ResourceController) UploadOne(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		controller.BadRequest(ctx, "文件不存在", nil)
		ctx.Abort()
		return
	}
	//判断上传资源的大小
	if file.Size > 1024*1024*10 {
		controller.BadRequest(ctx, "上传文件必须小于10MB", nil)
		ctx.Abort()
		return
	}
	//判断文件类型是否合法
	ext := path.Ext(file.Filename)
	if !utils.Contains(setting.Extension, ext) {
		controller.BadRequest(ctx, "文件格式不允许", nil)
		ctx.Abort()
		return
	}
	userid := uint(parseUid(ctx))
	//生成主键
	u, _ := uuid.NewUUID()
	uid := u.String()
	//生成文件的磁盘路径
	abpath := setting.ResourcePath + string(os.PathSeparator) + uid + ext
	if s := iresourceservice.CreateOne(uid, abpath, userid); s != "" {
		controller.InternalServerError(ctx, s, nil)
		ctx.Abort()
		return
	}
	if err := ctx.SaveUploadedFile(file, abpath); err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
		return
	}
	controller.Success(ctx, gin.H{
		"resource": "/api/resource/" + uid,
	})
}
