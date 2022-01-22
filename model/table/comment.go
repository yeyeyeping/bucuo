package table

type Comment struct {
	Model
	Content   string  `gorm:"type:varchar(3000);not null"`
	UserID    uint    `gorm:"not null"`
	LikeUsers []User  `gorm:"many2many:comment_like;constraint:OnDelete:CASCADE"`
	Replies   []Reply `gorm:"references:ID;constraint:OnDelete:CASCADE"`
	OwnerType string  `gorm:"type:varchar(20);not null"`
	OwnerID   uint    `gorm:"check:owner_id > 0"`
}
