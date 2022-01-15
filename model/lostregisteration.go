package model

import "gorm.io/gorm"

type LostRegisteration struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null"`
	Sno      string `gorm:"type:varchar(20);"`
	Location string `gorm:"type:varchar(100);not null"`
	Desc     string `gorm:"type:varchar(3000);"`
	Status   string `gorm:"type:varchar(20);"`
}
