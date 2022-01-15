package model

import "gorm.io/gorm"

/*
用户可以评论帖子
回复只能恢复评论
*/
type Reply struct {
	gorm.Model
	CommentID uint   `gorm:"not null"`
	Content   string `gorm:"type:varchar(3000);not null"`
	User      User
	UserID    uint   `gorm:"not null"`
	LikeUsers []User `gorm:"many2many:reply_like"`
}
