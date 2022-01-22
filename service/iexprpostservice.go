package service

import (
	"bucuo/model/request"
	"bucuo/model/response"
)

type IExprPostService interface {
	PushPost(req *request.ExprPostReq) string
	FindAll(column string, pagesize uint, pagenum uint) (result *[]response.SimpleExprPost, err error)
	FindDetails(id uint) (error, *response.SimpleExprDetailResp)
	DeleteOne(id uint) error
	ExprExist(exprid uint, uid uint) error
	UpdateOne(req *request.UpdateExprReq) error
}
