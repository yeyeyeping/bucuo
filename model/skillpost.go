package model

import "gorm.io/gorm"

type SkillPost struct {
	gorm.Model
	Title     string     `gorm:"type:varchar(20);not null"`
	Content   string     `gorm:"type:varchar(3000);not null"`
	Column    string     `gorm:"type:enum('交换','分享','求助');default:'求助'"`
	Labels    []Label    `gorm:"polymorphic:Owner;"`
	Comments  []Comment  `gorm:"polymorphic:Owner;"`
	Resources []Resource `gorm:"polymorphic:Owner;"`
	OwnerType string     `gorm:"type:varchar(20);not null"`
	OwnerID   uint       `gorm:"not null"`
}
