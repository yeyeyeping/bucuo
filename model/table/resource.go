package table

import (
	_ "github.com/google/uuid"
)

type Resource struct {
	ID           string `gorm:"default:uuid()"`
	DiskFilePath string `gorm:"type:varchar(500);not null"`
	UploaderID   uint   `gorm:"not null"`
	Uploader     *User  `gorm:"foreignKey:UploaderID"`
	OwnerType    string `gorm:"type:varchar(20)"`
	OwnerID      uint
}

//func (r *Resource) BeforeCreate(tx *gorm.DB) (err error) {
//	// UUID version 4
//	r.ID = uuid.NewString()
//	return
//}
