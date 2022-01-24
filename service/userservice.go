package service

import (
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/response"
	"bucuo/model/table"
	"math"
)

type UserService struct{}

var userdao dao.UserDao

func (u UserService) LoginByUserName(username string, password string) (string, bool) {
	ok, err := userdao.CheckLogin(username, password)
	if ok {
		return err, true
	} else {
		return "登录失败，用户名或密码错误", false
	}
}

func (u UserService) ValidateOpenId(s string) (bool, string) {
	return userdao.CheckOpenid(s)
}

func (u UserService) CreateUser(req *request.UserCreateReq) (bool, string) {
	err := userdao.CreateUser(&table.User{
		Sno:      req.Sno,
		Password: req.Sno,
		OpenId:   req.OpenId,
		Grade:    req.Grade,
		Username: req.Username,
		College:  req.College,
		Phone:    req.Phone,
	})
	if err != nil {
		return false, err.Error()
	} else {
		return true, ""
	}
}

func (u UserService) GetUserById(userid uint) *response.UserGetResp {
	return userdao.GetUserBy(userid)
}
func (u UserService) GetOtherUserById(userid uint, otherid uint) interface{} {
	if !userdao.UserExist(otherid) {
		return ""
	}
	return struct {
		*response.UserGetResp
		Like bool `json:"like"`
	}{
		UserGetResp: userdao.GetUserBy(otherid),
		Like:        userdao.LikeExist(userid, otherid),
	}
}

func (u UserService) UpdateUser(userid uint, req *request.UserCreateReq) error {
	return userdao.UpdateUser(userid, req)
}

func (u UserService) AddHomeLike(like uint, liked uint) string {
	return userdao.AddHomeLike(like, liked)
}

func (u UserService) DeleteHomeLike(like uint, liked uint) string {
	return userdao.DeleteHomeLike(like, liked)
}
func (u UserService) GetUserExpr(id uint) (*[]response.SimpleExprPost, error) {
	res, err := exprdao.GetUserExpr(id)
	if err != nil {
		return nil, err
	}
	return BuildSimpleExprPost(res), nil
}
func (u UserService) GetDetail(id uint) interface{} {
	num := exprdao.GetExprNum(id)
	return &struct {
		ExprPostNum  uint `json:"exprpostnum"`
		CollectorNum uint `json:"collectornum"`
		TimesOnList  uint
	}{
		num,
		exprdao.GetCollectors(id),
		uint(math.Round(float64(num / 10))),
	}
}
