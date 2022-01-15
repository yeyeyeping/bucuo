package model

import "gorm.io/gorm"

type Label struct {
	Content   string `gorm:"type:varchar(20);not null"`
	OwnerType string `gorm:"type:varchar(20);not null"`
	OwnerID   uint   `gorm:"not null"`
	gorm.Model
}
