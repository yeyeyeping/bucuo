package table

import (
	"gorm.io/gorm"
	"os"
)

type LocalPost struct {
	Title       string      `gorm:"type:varchar(12);not null"`
	Content     string      `gorm:"type:varchar(3000);not null"`
	Column      string      `gorm:"type:enum('美食','娱乐','生活购物');default:'娱乐'"`
	Labels      *[]Label    `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Comments    *[]Comment  `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Resources   *[]Resource `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	PublisherID uint        `gorm:"not null" json:"publisherID,omitempty"`
	Publisher   *User       `gorm:"foreignKey:PublisherID"`
	Model
}
type CommonPost struct {
	Title       string
	Content     string
	Column      string
	Labels      *[]Label    `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Comments    *[]Comment  `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Resources   *[]Resource `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	PublisherID uint        `gorm:"not null" json:"publisherID,omitempty"`
	ID          uint
}

func (e LocalPost) BeforeDelete(db *gorm.DB) error {
	v := LocalPost{Model: Model{ID: e.ID}}
	db.
		Table("local_posts").
		Preload("Resources").
		First(&v)
	if v.Labels != nil {
		db.Table("labels").Unscoped().Delete(v.Labels, "owner_id", v.ID)
	}
	if v.Resources != nil {
		for _, resource := range *v.Resources {
			os.Remove(resource.DiskFilePath)
		}
		db.Table("resources").Unscoped().Delete(v.Resources, "owner_id", v.ID)
	}
	return nil
}
