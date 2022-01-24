package table

// Reply /*
type Reply struct {
	Model
	CommentID     uint    `gorm:"not null"`
	TargetComment Comment `gorm:"foreignKey:CommentID"`
	Content       string  `gorm:"type:varchar(300);not null"`
	ReplierID     uint    `gorm:"not null"`
	Replier       *User   `gorm:"foreignKey:ReplierID"`
	LikeUsers     []User  `gorm:"many2many:reply_like;constraint:OnDelete:CASCADE"`
}
