package request

type UserReq struct {
	UserName string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `json:"password" validate:"required,min=6,max=16" label:"密码"`
}

type UserCreateReq struct {
	Sno            string `json:"sno" validate:"required,max=10" label:"学号"`
	Grade          string `json:"grade" validate:"required,max=6" label:"年级"`
	College        string `json:"college" validate:"required" label:"学院"`
	Phone          string `json:"phone" validate:"required" label:"手机号"`
	ProfilePicture string `json:"profilepicture" validate:"required" label:"头像链接"`
	OpenId         string `json:"-"`
}

type ExprPostReq struct {
	Title     string   `json:"title" validate:"required,max=10" label:"标题"`
	Content   string   `json:"content" validate:"required,min=10" label:"内容"`
	Column    string   `json:"column" validate:"required,oneof='课程考试' '考研保研' '竞赛考证' '新生守则' '其他经验'" label:"分类"`
	Labels    []string `json:"labels" validate:"required,min=1,max=4,dive,max=10" label:"标签"`
	PublishId uint     `json:"-" validate:"-"`
}

type ByPage struct {
	PageSize uint   `json:"pagesize" validate:"required,min=1,max=10" label:"页大小"`
	PageNum  uint   `json:"pagenum" validate:"required,min=1" label:"页序号"`
	Column   string `json:"column" validate:"oneof='课程考试' '考研保研' '竞赛考证' '新生守则' '其他经验' ''" label:"分类"`
}

type IGetOwnerInfo interface {
	GetOwnerType() string
	GetOwnerId() uint
	GetUserID() uint
}
type ResourceReq struct {
	OwnerID   uint   `json:"ownerID" validate:"required,min=1"`
	OwnerType string `json:"ownerType" validate:"required,min=1"`
	UserID    uint   `validate:"-"`
}

type UpdateExprReq struct {
	ID      uint   `json:"id" validate:"required,min=1" label:"编号"`
	Title   string `json:"title" validate:"min=3,max=10" label:"标题"`
	Content string `json:"content" validate:"min=10" label:"内容"`
	Column  string `json:"column" validate:"oneof='课程考试' '考研保研' '竞赛考证' '新生守则' '其他经验'" label:"分类"`
}

func (r ResourceReq) GetUserID() uint {
	return r.UserID
}
func (r ResourceReq) GetOwnerId() uint {
	return r.OwnerID
}
func (r ResourceReq) GetOwnerType() string {
	return r.OwnerType
}
