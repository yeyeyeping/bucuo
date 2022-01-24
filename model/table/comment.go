package table

import (
	"gorm.io/gorm"
)

type Comment struct {
	Model
	Content   string  `gorm:"type:varchar(600);not null"`
	UserID    uint    `gorm:"not null"`
	User      User    `gorm:"foreignKey:UserID"`
	LikeUsers []User  `gorm:"many2many:comment_like;constraint:OnDelete:CASCADE"`
	Replies   []Reply `gorm:"references:ID;constraint:OnDelete:CASCADE"`
	OwnerType string  `gorm:"type:varchar(20);not null"`
	OwnerID   uint    `gorm:"check:owner_id > 0"`
}

func (c Comment) BeforeDelete(db *gorm.DB) error {
	db.Where("CommentID=?", c.ID).Delete(Reply{})
	return nil
}
