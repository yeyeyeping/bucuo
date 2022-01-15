package model

import "gorm.io/gorm"

type LocalPost struct {
	Title     string     `gorm:"type:varchar(12);not null"`
	Content   string     `gorm:"type:varchar(3000);not null"`
	Column    string     `gorm:"type:enum('美食','娱乐','生活购物');default:'娱乐'"`
	Labels    []Label    `gorm:"polymorphic:Owner;"`
	Comments  []Comment  `gorm:"polymorphic:Owner;"`
	Resources []Resource `gorm:"polymorphic:Owner;"`
	OwnerType string     `gorm:"type:varchar(20);not null"`
	OwnerID   uint       `gorm:"not null"`
	gorm.Model
}
