package dao

import (
	"bucuo/model/request"
	"bucuo/model/table"
	"gorm.io/gorm"
)

type ExprDao struct {
}

func (ExprDao) PostExpr(post *table.ExprPost) (es string) {
	if err := DB.Model(&table.ExprPost{}).Create(post).Error; err != nil {
		return err.Error()
	}
	return ""
}
func (ExprDao) FindAll(column string, pagesize uint, pagenum uint) (postlist *[]table.ExprPost, err error) {
	if column == "" {
		err = DB.
			Table("expr_posts").
			Limit(int(pagesize)).
			Offset(int((pagenum-1)*pagesize)).
			Preload("Labels").
			Preload("Collectors", func(db *gorm.DB) *gorm.DB {
				return db.Select("ID")
			}).
			Find(&postlist).Error
	} else {
		err = DB.
			Table("expr_posts").
			Where("`column`=?", column).
			Limit(int(pagesize)).
			Offset(int((pagenum-1)*pagesize)).
			Preload("Labels").
			Preload("Collectors", func(db *gorm.DB) *gorm.DB {
				return db.Select("ID")
			}).
			Find(&postlist).Error
	}
	return
}
func (ExprDao) FindOne(id uint) (*table.ExprPost, error) {
	p := &table.ExprPost{
		Model: table.Model{ID: id},
	}
	return p, DB.
		Table("expr_posts").
		Preload("Publisher").
		Preload("Labels").
		Preload("Comments").
		Preload("Collectors").
		First(p).Error
}
func (ExprDao) DeleteOne(id uint) error {
	return DB.
		Unscoped().
		Table("expr_posts").
		Delete(table.ExprPost{Model: table.Model{ID: id}}).Error
}
func (ExprDao) ExistOne(exprid uint, uid uint) error {
	return DB.Table("expr_posts").Where("id=? and publisher_id=?", exprid, uid).First(&table.ExprPost{}).Error
}
func (ExprDao) UpdateOne(req *request.UpdateExprReq) error {
	return DB.Table("expr_posts").Updates(req).Error
}
