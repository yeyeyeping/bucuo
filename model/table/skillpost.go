package table

import "gorm.io/gorm"

type SkillPost struct {
	Model
	Title       string      `gorm:"type:varchar(20);not null"`
	Content     string      `gorm:"type:varchar(3000);not null"`
	Column      string      `gorm:"type:enum('交换','分享','求助');default:'求助'"`
	Labels      *[]Label    `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Comments    *[]Comment  `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	Resources   *[]Resource `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE"`
	PublisherID uint        `gorm:"not null"`
	Publisher   *User       `gorm:"foreignKey:PublisherID" json:"publisher,omitempty"`
}

func (e SkillPost) BeforeDelete(db *gorm.DB) error {
	v := SkillPost{Model: Model{ID: e.ID}}
	db.
		Table("skill_posts").
		Preload("Labels").
		First(&v)
	if v.Labels != nil {
		db.Table("labels").Unscoped().Delete(v.Labels)
	}
	return nil
}
