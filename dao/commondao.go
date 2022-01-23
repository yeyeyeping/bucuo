package dao

import (
	"bucuo/model/table"
	"gorm.io/gorm"
)

type CommonDao struct {
}

func (d CommonDao) CreatePost(
	posttype string,
	title string,
	content string,
	column string,
	publisherid uint,
	lss *[]string,
) (error, uint) {
	t := &struct {
		ID          uint
		Content     string
		Title       string
		Column      string
		PublisherID uint
	}{
		Content:     content,
		Title:       title,
		Column:      column,
		PublisherID: publisherid,
	}
	if err := DB.Table(posttype).
		Create(t).Error; err != nil {
		return err, 0
	}
	var ls = make([]table.Label, len(*lss))
	for i, l := range *lss {
		ls[i] = table.Label{
			Content:   l,
			OwnerType: posttype,
			OwnerID:   t.ID,
		}
	}
	err = DB.Table("labels").Create(&ls).Error
	if err != nil {
		return err, 0
	}
	return nil, t.ID
}
func (d CommonDao) AddResource(posttype string, postid uint, resourcesid *[]string) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, rid := range *resourcesid {
			err := DB.Exec("update resources set owner_type=?,owner_id=? where id=?", posttype, postid, rid).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
func (d CommonDao) FindAllSkillPost(posttype string, column string, pagesize uint, pagenum uint) (*[]table.SkillPost, error) {
	rs := &[]table.SkillPost{}
	var er error
	if column != "" {
		er = DB.
			Table(posttype).
			Where("`column`=?", column).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Preload("Resources").
			Preload("Labels").
			Find(rs).Error
	} else {
		er = DB.
			Table(posttype).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Preload("Resources").
			Preload("Labels").
			Find(rs).Error
	}
	return rs, er
}
func (d CommonDao) FindAllLocalPost(posttype string, column string, pagesize uint, pagenum uint) (*[]table.LocalPost, error) {
	rs := &[]table.LocalPost{}
	var er error
	if column != "" {
		er = DB.
			Table(posttype).
			Where("`column`=?", column).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Preload("Resources").
			Preload("Labels").
			Find(rs).Error
	} else {
		er = DB.
			Table(posttype).
			Limit(int(pagesize)).
			Offset(int((pagenum - 1) * pagesize)).
			Preload("Resources").
			Preload("Labels").
			Find(rs).Error
	}
	return rs, er
}
func (d CommonDao) Exist(posttype string, ownerid uint, uid uint) bool {
	var i int64
	DB.Table(posttype).Where("id=? and publisher_id=?", ownerid, uid).Count(&i)
	return i == 1
}
func (d CommonDao) Delete(posttype string, id uint, uid uint) string {
	var i *gorm.DB
	if posttype == "local_posts" {
		i = DB.Table(posttype).Where("id=? and publisher_id=?", id, uid).Delete(&table.LocalPost{
			Model: table.Model{ID: id},
		})
	} else if posttype == "skill_posts" {
		i = DB.Table(posttype).Where("id=? and publisher_id=?", id, uid).Delete(&table.SkillPost{
			Model: table.Model{ID: id},
		})
	}
	if i.Error != nil {
		return i.Error.Error()
	}
	if i.RowsAffected == 0 {
		return "未找到记录"
	}
	if i.RowsAffected == 1 {
		return ""
	}
	return "未知错误"
}
func (d CommonDao) FindOneSkillPost(id uint) (*table.SkillPost, error) {
	rs := &table.SkillPost{
		Model: table.Model{ID: id},
	}
	err := DB.
		Table("skill_posts").
		Preload("Publisher").
		Preload("Labels").
		Preload("Comments").
		Preload("Resources").
		First(rs).Error
	return rs, err
}
func (d CommonDao) FindOneLocalPost(id uint) (*table.LocalPost, error) {
	rs := &table.LocalPost{Model: table.Model{ID: id}}
	err := DB.
		Table("local_posts").
		Preload("Publisher").
		Preload("Labels").
		Preload("Comments").
		Preload("Resources").
		First(rs).Error
	return rs, err
}
func (d CommonDao) Update(id uint, title string, content string, column string, publisher_id uint, tp string) string {
	tmpdb := DB.Exec("update `"+tp+"` "+"set title=?,content=?,`column`=? where id=? and publisher_id=?",
		title, content, column, id, publisher_id)
	if tmpdb.RowsAffected == 0 {
		return "找不到记录"
	}
	if tmpdb.Error != nil {
		return tmpdb.Error.Error()
	}
	return ""
}
func (d CommonDao) ClearLabels(ownertype string, ownerid uint) string {
	tmpdb := DB.Exec("delete from labels where owner_type=? and owner_id=?", ownertype, ownerid)
	if tmpdb.RowsAffected == 0 {
		return "找不到记录"
	}
	if tmpdb.Error != nil {
		return tmpdb.Error.Error()
	}
	return ""
}
func (d CommonDao) CreateLabels(ownertype string, ownerid uint, label *[]table.Label) error {
	if ownertype == "local_posts" {
		return DB.Model(&table.LocalPost{Model: table.Model{ID: ownerid}}).Association("Labels").Append(label)
	}
	if ownertype == "skill_posts" {
		return DB.Model(&table.SkillPost{Model: table.Model{ID: ownerid}}).Association("Labels").Append(label)
	}
	return gorm.ErrUnsupportedRelation
}
func (d CommonDao) ReplaceResources(ownertype string, ownerid uint, resource *[]table.Resource) error {
	if ownertype == "local_posts" {
		return DB.Model(&table.LocalPost{Model: table.Model{ID: ownerid}}).Association("Resources").Replace(resource)
	}
	if ownertype == "skill_posts" {
		return DB.Model(&table.SkillPost{Model: table.Model{ID: ownerid}}).Association("Resources").Replace(resource)
	}
	return gorm.ErrUnsupportedRelation

}
