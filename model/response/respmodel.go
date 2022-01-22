package response

import (
	"bucuo/constant/errormsg"
	"time"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

type BaseResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespModel(code int, msg string, data interface{}) interface{} {
	return BaseResp{
		Code: code,
		Msg:  errormsg.GetMsg(code, msg),
		Data: data,
	}
}

type UserGetResp struct {
	Sno               string
	Grade             string
	College           string
	BucuoId           string
	KnowledgeCurrency uint
	IsExprVip         bool
	IsStarVip         bool
	HomeLikeNum       int64
	VipExpr           uint
	VipStar           uint
}
type SimpleLabel struct {
	LabelID uint   `json:"label_id"`
	Content string `json:"content"`
}
type SimpleExprPost struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Content      string         `json:"content"`
	Column       string         `json:"column"`
	Labels       *[]SimpleLabel `json:"labels"`
	CollectorNum uint
}
type SimpleExprDetailResp struct {
	ID             uint           `json:"id"`
	BucuoId        string         `json:"BucuoId"`
	PublisherTime  Time           `json:"publisherTime"`
	ProfilePicture string         `json:"profilePicture"`
	Labels         *[]SimpleLabel `json:"labels"`
	CollectorNum   int            `json:"collectorNum"`
	CommentNum     int            `json:"commentNum"`
}
