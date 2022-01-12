package main

import (
	"bucuo/dao"
	"bucuo/model"
	"bucuo/service"
	"log"
)

func main() {
	dao.Ok()
	var userList []model.User
	service.UserService{}.FindAll(&userList)
	log.Printf("%#v", userList)
}
