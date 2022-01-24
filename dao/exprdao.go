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
			Order("id").
			Offset(int((pagenum-1)*pagesize)).
			Limit(int(pagesize)).
			Preload("Labels").
			Preload("Comments", func(db *gorm.DB) *gorm.DB {
				return db.Select("ID")
			}).
			Preload("Collectors", func(db *gorm.DB) *gorm.DB {
				return db.Select("ID")
			}).
			Find(&postlist).Error
	} else {
		err = DB.
			Table("expr_posts").
			Where("`column`=?", column).
			Order("id").
			Offset(int((pagenum-1)*pagesize)).
			Limit(int(pagesize)).
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
func (ExprDao) GetUserExpr(uid uint) (*[]table.ExprPost, error) {
	res := make([]table.ExprPost, 0)
	err := DB.Table("expr_posts").
		Preload("Labels").
		Preload("Collectors", func(db *gorm.DB) *gorm.DB {
			return db.Select("ID")
		}).
		Where("publisher_id=?", uid).
		Find(&res).Error
	return &res, err
}
func (ExprDao) GetExprNum(uid uint) uint {
	i := 0
	DB.Raw("select count(*) from expr_posts where publisher_id=?", uid).Scan(&i)
	return uint(i)
}
func (ExprDao) GetCollectors(uid uint) uint {
	i := 0
	DB.Raw(
		"select count(*) from expr_post_collect_collected where expr_post_id in "+
			"(select id from expr_posts where publisher_id=?)", uid).
		Scan(&i)
	return uint(i)
}
