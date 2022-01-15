package dao

import "bucuo/model"

type UserDao struct{}

func (dao UserDao) GetUser(id int) (user model.User) {
	DB.First(&user, id)
	return
}
func (dao UserDao) CreateUser(user *model.User) {
	DB.Create(user)
}

//dao.UserDao{}.CreateUser(&model.User{
//	Sno:      "1904060087",
//	Password: "yeep2020",
//	Username: "yeep",
//	OpenId:   "123",
//	College:  "河南大学",
//	Phone:    "18712337064",
//})
//dao.DB.Create(&model.ExprPost{
//	Title:      "测试",
//	Content:    "测试测试",
//	Column:     exprpostcolumn.COMPETITION,
//	Labels:     []model.Label{{Content: "商学院"}},
//	Replies:    nil,
//	OwnerType:  "",
//	OwnerID:    0,
//	Collectors: nil,
//	Model:      gorm.Model{},
//})
