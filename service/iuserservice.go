package service

import (
	"bucuo/model/request"
	"bucuo/model/response"
)

type IUserServiece interface {
	//RegisterUser(user user.User) error
	LoginByUserName(string, string) (string, bool)
	ValidateOpenId(string) (bool, string)
	CreateUser(req *request.UserCreateReq) (bool, string)
	GetUserById(userid uint) *response.UserGetResp
	UpdateUser(userid uint, req *request.UserCreateReq) error
	AddHomeLike(like uint, liked uint) string
	DeleteHomeLike(like uint, liked uint) string
	GetOtherUserById(userid uint, otherid uint) interface{}
	GetUserExpr(id uint) (*[]response.SimpleExprPost, error)
	GetDetail(id uint) interface{}
}
