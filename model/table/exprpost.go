package table

import "gorm.io/gorm"

type ExprPost struct {
	Title       string     `gorm:"type:varchar(20);not null"`
	Content     string     `gorm:"type:varchar(3000);not null"`
	Column      string     `gorm:"type:enum('课程考试','考研保研','竞赛考证','新生守则','其他经验');default:'其他经验'"`
	Labels      *[]Label   `gorm:"polymorphic:Owner;constraint:OnDelete:CASCADE" json:"labels,omitempty"`
	Comments    *[]Comment `gorm:"polymorphic:Owner; constraint:OnDelete:CASCADE" json:"comments,omitempty"`
	PublisherID uint       `gorm:"not null"`
	Publisher   *User      `gorm:"foreignKey:PublisherID" json:"publisher,omitempty"`
	Collectors  *[]User    `gorm:"many2many:expr_post_collect_collected; constraint:OnDelete:CASCADE" json:"collectors,omitempty"`
	Model
}

func (e ExprPost) BeforeDelete(db *gorm.DB) error {
	v := ExprPost{Model: Model{ID: e.ID}}
	db.
		Table("expr_posts").
		First(&v)
	if v.Labels != nil {
		db.Table("labels").Unscoped().Delete(v.Labels, "owner_id=?", v.ID)
	}
	return nil
}
