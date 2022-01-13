package service

import (
	"bucuo/dao"
	"bucuo/model"
)

type UserService struct{}

func (u UserService) AddUser(user *model.User) error {
	return dao.DB.Create(user).Error
}
func (u UserService) DeleteUser(id int) error {
	return dao.DB.Delete(&model.User{}, id).Error
}
func (u UserService) UpdateUser(user *model.User) error {
	return dao.DB.Save(user).Error
}
func (u UserService) FindById(id int, user *model.User) error {
	return dao.DB.Find(user, id).Error
}
func (u UserService) FindAll(userList *[]model.User) error {
	return dao.DB.Find(userList).Error
}
func (u UserService) FindByUserName(user *model.User) bool {
	return dao.DB.Find(user, user).RowsAffected == 1
}
