package model

import (
	"bucuo/util"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	Sno                 string     `gorm:"type:varchar(20);not null"`
	Password            string     `json:"password" form:"password" uri:"password" gorm:"type:varchar(32)"`
	Username            string     `json:"username" form:"username" uri:"username"`
	OpenId              string     `gorm:"type:varchar(100);not null"`
	College             string     `gorm:"type:varchar(32);not null"`
	Phone               string     `gorm:"type:varchar(22);not null"`
	BucuoID             string     `gorm:"varchar();not null"`
	PublishExprPosts    []ExprPost `gorm:"polymorphic:Owner;"`
	CollectionExprPosts []ExprPost `gorm:"many2many:user_exprpost;References:id"`
	HomePageLikeUserId  uint
	HomePageLikeUser    []User `gorm:"many2many:user_home_page_like_user"`
	VipExpr             uint
	KnowledgeCurrency   uint
	IsExprVip           bool `gorm:"DEFAULT:false"`
	IsStarVip           bool `gorm:"DEFAULT:false"`
	FriendID            uint
	Friends             []User `gorm:"many2many:user_friend;"`
	gorm.Model
}

func (r *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(r).Update("bucuo_id", util.IdPrefix+strconv.FormatUint(uint64(r.ID), 10))
	return
}
