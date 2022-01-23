package dao

import (
	"bucuo/model/table"
	"gorm.io/gorm"
)

type CommentDao struct {
}

func (d CommentDao) AddComment(comment *table.Comment) error {
	return DB.Table("comments").Create(comment).Error
}
func (d CommentDao) DeleteComment(id uint, uid uint) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if DB.Table("comments").
			Where("id=? and user_id=?", id, uid).
			Delete(&table.Comment{Model: table.Model{ID: id}}).
			RowsAffected == 1 {
			return nil
		} else {
			return gorm.ErrRecordNotFound
		}
	})
}
func (d CommentDao) AddReply(commentid uint, reply *table.Reply) error {
	return DB.Model(&table.Comment{
		Model: table.Model{ID: commentid},
	}).Association("Replies").Append(reply)
}
func (d CommentDao) DeleteReply(id uint, uid uint) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if DB.Table("replies").Where("id=? and uid=?", id, uid).Delete(table.Reply{}).RowsAffected == 1 {
			return nil
		} else {
			return gorm.ErrRecordNotFound
		}
	})
}
