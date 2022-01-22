package dao

import (
	"bucuo/model/request"
	"bucuo/model/response"
	"bucuo/model/table"
	"strconv"
)

type UserDao struct{}

func (dao UserDao) CheckLogin(username string, password string) (bool, string) {
	user := table.User{}
	err = DB.First(&user, "username=?", username).Error
	if err != nil {
		return false, err.Error()
	}
	return password == user.Password, strconv.FormatUint(uint64(user.ID), 10)
}
func (u UserDao) CheckOpenid(openid string) (bool, string) {
	user := table.User{}
	if err := DB.First(&user, "open_id=?", openid).Error; err != nil {
		return false, err.Error()
	}
	return true, strconv.FormatUint(uint64(user.ID), 10)
}
func (u UserDao) GetUserBy(userid uint) (user *response.UserGetResp) {
	usermodel := &table.User{Model: table.Model{ID: userid}}
	DB.Model(usermodel).
		Select("Sno", "Grade", "College", "BucuoID", "IsExprVip",
			"IsStarVip", "VipExpr", "VipStar", "KnowledgeCurrency").
		Scan(&user)
	DB.
		Raw("SELECT COUNT(*) FROM home_page_like_liked where like_home_page_user_id=?",
			userid).
		Row().
		Scan(&user.HomeLikeNum)
	return
}
func (u UserDao) CreateUser(user *table.User) error {
	return DB.Create(user).Error
}
func (u UserDao) UpdateUser(userid uint, user *request.UserCreateReq) error {
	if err := DB.First(&table.User{
		Model: table.Model{
			ID: userid,
		},
	}).Omit("OpenId").Updates(user).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (u UserDao) AddHomeLike(like uint, liked uint) string {
	if err := DB.First(&table.User{Model: table.Model{ID: liked}}).Error; err != nil {
		return err.Error()
	}
	var i int64
	DB.Raw("SELECT COUNT(*) FROM home_page_like_liked where like_home_page_user_id=? and user_id=?", liked, like).Scan(&i)
	if i != 0 {
		return "您已经点过赞了！"
	}
	DB.Model(&table.User{
		Model: table.Model{ID: like},
	}).
		Association("LikeHomePageUsers").
		Append(&table.User{Model: table.Model{ID: liked}})
	return ""
}
func (u UserDao) DeleteHomeLike(like uint, liked uint) string {
	if err := DB.First(&table.User{Model: table.Model{ID: liked}}).Error; err != nil {
		return err.Error()
	}
	var i int64
	DB.Raw("SELECT COUNT(*) FROM home_page_like_liked where like_home_page_user_id=? and user_id=?", liked, like).Scan(&i)
	if i == 0 {
		return "您还没有为他点赞哦！"
	}
	DB.Model(&table.User{
		Model: table.Model{ID: like},
	}).
		Association("LikeHomePageUsers").
		Delete(&table.User{Model: table.Model{ID: liked}})
	return ""
}
