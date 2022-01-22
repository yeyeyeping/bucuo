package dao

import (
	"bucuo/model/request"
	"bucuo/model/table"
)

type ResourceDao struct {
}

func (r ResourceDao) Upload(post *table.Resource) error {
	return DB.
		Model(&table.Resource{}).
		Create(post).Error
}
func (r ResourceDao) PostExist(post request.IGetOwnerInfo) (bool, error) {
	var i int64
	err := DB.
		Table(post.GetOwnerType()).
		Where("id=? and publisher_id=?", post.GetOwnerId(), post.GetUserID()).
		Count(&i).Error
	return i == 1, err
}
func (r ResourceDao) GetPostCount(post request.IGetOwnerInfo) (value int64, err error) {
	err = DB.
		Table("resources").
		Where("owner_id=? and owner_type=? and uploader_id=?", post.GetOwnerId(), post.GetOwnerType(), post.GetUserID()).
		Count(&value).Error
	return
}
func (r ResourceDao) GetResouce(uuid string) (string, error) {
	res := &table.Resource{ID: uuid}
	if err := DB.Table("resources").First(res).Error; err != nil {
		return "", err
	} else {
		return res.DiskFilePath, nil
	}

}
func (r ResourceDao) CreateOne(resource *table.Resource) error {
	return DB.Create(resource).Error
}
