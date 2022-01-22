package service

import (
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/table"
)

type ResourceService struct {
}

var resourcedao dao.ResourceDao

func (r ResourceService) PostExist(req *request.ResourceReq) (bool, error) {
	return resourcedao.PostExist(req)
}

func (r ResourceService) GetPostCount(post request.IGetOwnerInfo) (value int64, err error) {
	return resourcedao.GetPostCount(post)
}

func (r ResourceService) CreateResource(ID string, DiskFilePath string, UploaderID uint, OwnerType string, OwnerID uint) error {
	rs := &table.Resource{
		ID:           ID,
		DiskFilePath: DiskFilePath,
		UploaderID:   UploaderID,
		OwnerType:    OwnerType,
		OwnerID:      OwnerID,
	}
	return dao.DB.
		Table("resources").
		Create(rs).Error

}

func (r ResourceService) GetResouce(uuid string) (string, error) {
	return resourcedao.GetResouce(uuid)
}
func (r ResourceService) CreateOne(id string, abpath string, uid uint) string {
	rs := &table.Resource{
		ID:           id,
		DiskFilePath: abpath,
		UploaderID:   uid,
	}
	if err := resourcedao.CreateOne(rs); err != nil {
		return err.Error()
	}
	return ""
}
