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
		if DB.Table("replies").Where("id=? and replier_id=?", id, uid).Delete(table.Reply{}).RowsAffected == 1 {
			return nil
		} else {
			return gorm.ErrRecordNotFound
		}
	})
}
func (c CommentDao) PostExist(t string, id uint) bool {
	i := 0
	DB.Raw("select count(*) from "+t+" where id =?", id).Scan(&i)
	return i == 1
}

type ExistErr struct {
}

func (e ExistErr) Error() string {
	return "记录已存在"
}

func (d CommentDao) Like(id uint, commentid uint) error {
	if !d.CommentExist(commentid) {
		return gorm.ErrRecordNotFound
	}
	if d.LikeExist(id, commentid) {
		return ExistErr{}
	}
	comment := &table.Comment{Model: table.Model{ID: commentid}}
	err := DB.Model(comment).
		Association("LikeUsers").
		Append(&table.User{Model: table.Model{ID: id}})
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (d CommentDao) LikeReply(id uint, replyid uint) error {
	if !d.ReplyExist(replyid) {
		return gorm.ErrRecordNotFound
	}
	if d.LikeReplyExist(id, replyid) {
		return ExistErr{}
	}
	err := DB.Model(&table.Reply{Model: table.Model{ID: replyid}}).
		Association("LikeUsers").
		Append(&table.User{Model: table.Model{ID: id}})
	if err != nil {
		return err
	}
	return nil
}
func (d CommentDao) ReplyExist(id uint) bool {
	i := 0
	DB.Raw("select count(*) from replies where id=?", id).Scan(&i)
	return i == 1
}
func (d CommentDao) CommentExist(id uint) bool {
	i := 0
	DB.Raw("select count(*) from comments where id=?", id).Scan(&i)
	return i == 1
}
func (d CommentDao) LikeExist(uid uint, comentid uint) bool {
	i := 0
	DB.Raw("select count(*) from comment_like where comment_id=? and user_id=?", comentid, uid).Scan(&i)
	return i == 1
}
func (d CommentDao) LikeReplyExist(uid uint, replyid uint) bool {
	i := 0
	DB.Raw("select count(*) from reply_like where reply_id=? and user_id=?", replyid, uid).Scan(&i)
	return i == 1
}
