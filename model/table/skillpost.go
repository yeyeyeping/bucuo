package table

import (
	"gorm.io/gorm"
	"os"
)

type SkillPost struct {
	Title       string      `gorm:"type:varchar(20);not null"`
	Content     string      `gorm:"type:varchar(3000);not null"`
	Column      string      `gorm:"type:enum('交换','分享','求助');default:'求助'"`
	Labels      *[]Label    `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Comments    *[]Comment  `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Resources   *[]Resource `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	PublisherID uint        `gorm:"not null"`
	Publisher   *User       `gorm:"foreignKey:PublisherID" json:"publisher,omitempty"`
	Model
}

func (e SkillPost) BeforeDelete(db *gorm.DB) error {
	v := SkillPost{Model: Model{ID: e.ID}}
	db.
		Table("skill_posts").
		Preload("Resources").
		First(&v)
	if v.Labels != nil {
		db.Table("labels").Unscoped().Delete(v.Labels, "owner_id=?", v.ID)
	}
	if v.Resources != nil {
		for _, i2 := range *v.Resources {
			os.Remove(i2.DiskFilePath)
		}
		db.Table("resources").Unscoped().Delete(v.Resources, "owner_id=?", v.ID)
	}

	return nil
}
