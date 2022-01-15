package model

import (
	"gorm.io/gorm"
)

type ExprPost struct {
	Title      string    `gorm:"type:varchar(20);not null"`
	Content    string    `gorm:"type:varchar(3000);not null"`
	Column     string    `gorm:"type:enum('课程考试','考研保研','竞赛考证','新生守则','其他经验');default:'其他经验'"`
	Labels     []Label   `gorm:"polymorphic:Owner;"`
	Comments   []Comment `gorm:"polymorphic:Owner;"`
	OwnerType  string    `gorm:"type:varchar(20);not null"`
	OwnerID    uint      `gorm:"check:owner_id > 0"`
	Collectors []*User   `gorm:"many2many:user_expr_post;References:id"`
	gorm.Model
}
