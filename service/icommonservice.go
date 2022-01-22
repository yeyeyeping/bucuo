package service

import (
	"bucuo/model/request"
	"bucuo/model/response"
)

type ICommonService interface {
	PushPost(req *request.CommonPostReq) string
	FindAll(posttype string, column string, pagesize uint, pagenum uint) (string, *[]response.SimpleCommonPost)
	Exist(req *request.DeleteCommonReq, uid uint) bool
	Delete(posttype string, id uint, uid uint) string
	FindDetail(req *request.DeleteCommonReq) (*response.SimpleCommonPostDetailResp, string)
}
