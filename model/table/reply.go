package table

/*
用户可以评论帖子
回复只能恢复评论
*/
type Reply struct {
	Model
	CommentID uint    `gorm:"not null"`
	Content   string  `gorm:"type:varchar(3000);not null"`
	ReplierID uint    `gorm:"not null"`
	Replier   *User   `gorm:"foreignKey:ReplierID"`
	LikeUsers *[]User `gorm:"many2many:reply_like;constraint:OnDelete:CASCADE"`
}
