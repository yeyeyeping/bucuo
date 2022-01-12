package service

import (
	"bucuo/dao"
	"bucuo/model"
)

type UserService struct{}

func (u UserService) AddUser(user *model.User) (err error) {
	err = dao.DB.Create(user).Error
	return
}
func (u UserService) DeleteUser(id int) (err error) {
	err = dao.DB.Delete(&model.User{}, id).Error
	return
}
func (u UserService) UpdateUser(user *model.User) (err error) {
	err = dao.DB.Save(user).Error
	return
}
func (u UserService) FindById(id int, user *model.User) (err error) {
	err = dao.DB.Find(user, id).Error
	return
}
func (u UserService) FindAll(userList *[]model.User) (err error) {
	err = dao.DB.Find(userList).Error
	return
}
