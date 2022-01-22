package table

import (
	"bucuo/util/setting"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	Sno                 string      `gorm:"type:varchar(20);not null"`
	Password            string      `json:"password" form:"password" uri:"password" gorm:"type:varchar(32)"`
	Username            string      `json:"username" form:"username" uri:"username"`
	OpenId              string      `gorm:"type:varchar(100);not null"`
	College             string      `gorm:"type:varchar(32);not null"`
	Grade               string      `gorm:"type:varchar(20);not null"`
	Phone               string      `gorm:"type:varchar(22);not null"`
	BucuoID             string      `gorm:"varchar(20);not null"`
	PublishExprPosts    *[]ExprPost `gorm:"references:ID;foreignKey:PublisherID"`
	CollectionExprPosts *[]ExprPost `gorm:"many2many:expr_post_collect_collected"`
	LikeHomePageUsers   *[]User     `gorm:"many2many:home_page_like_liked"`
	ProfilePicture      string      `gorm:"type:varchar(300);not null"`
	VipExpr             uint
	VipStar             uint
	KnowledgeCurrency   uint
	IsExprVip           bool
	IsStarVip           bool
	Friends             *[]User `gorm:"many2many:user_friend;"`
	Model
}

func (r *User) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(r).Update("bucuo_id", setting.IdPrefix+strconv.FormatUint(uint64(r.ID), 10))
	return
}
