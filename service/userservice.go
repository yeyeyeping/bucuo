package service

import (
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/response"
	"bucuo/model/table"
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
	return err, false
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
		Username: req.Sno,
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

func (u UserService) UpdateUser(userid uint, req *request.UserCreateReq) error {
	return userdao.UpdateUser(userid, req)
}

func (u UserService) AddHomeLike(like uint, liked uint) string {
	return userdao.AddHomeLike(like, liked)
}

func (u UserService) DeleteHomeLike(like uint, liked uint) string {
	return userdao.DeleteHomeLike(like, liked)
}

//func (u UserService) AddUser(user *table.Model.User) error {
//	return dao.DB.Create(user).Error
//}
//func (u UserService) DeleteUser(id int) error {
//	return dao.DB.Delete(&table.Model.User{}, id).Error
//}
//func (u UserService) UpdateUser(user *table.Model.User) error {
//	return dao.DB.Save(user).Error
//}
//func (u UserService) FindById(id int, user *table.Model.User) error {
//	return dao.DB.Find(user, id).Error
//}
//func (u UserService) FindAll(userList *[]table.Model.User) error {
//	return dao.DB.Find(userList).Error
//}
//func (u UserService) FindByUserName(user *table.Model.User) bool {
//	return dao.DB.Find(user, user).RowsAffected == 1
//}
