package service

import (
	"bucuo/model/request"
)

type IResourceService interface {
	PostExist(req *request.ResourceReq) (bool, error)
	GetPostCount(post request.IGetOwnerInfo) (value int64, err error)
	CreateResource(ID string, DiskFilePath string, UploaderID uint, OwnerType string, OwnerID uint) error
	GetResouce(uuid string) (string, error)
	CreateOne(id string, abpath string, uid uint) string
}
