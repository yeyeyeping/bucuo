package dao

import "bucuo/model"

type UserDao struct{}

func (dao UserDao) GetUser(id int) (user model.User) {
	DB.First(&user, id)
	return
}
func (dao UserDao) CreateUser(user model.User) {
	DB.Create(&user)
}
