package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:varchar(3000);not null"`
	UserID    uint   `gorm:"not null"`
	User      User
	LikeUsers []User `gorm:"many2many:comment_like"`
	Replies   []Reply
	OwnerType string `gorm:"type:varchar(20);not null"`
	OwnerID   uint   `gorm:"not null"`
}
