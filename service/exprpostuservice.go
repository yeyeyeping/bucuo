package service

import (
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/response"
	"bucuo/model/table"
)

type ExprPostService struct{}

var exprdao dao.ExprDao

func BuildSimpleExprPost(result *[]table.ExprPost) (rs *[]response.SimpleExprPost) {
	rsp := make([]response.SimpleExprPost, len(*result))
	for i, p := range *result {
		ls := make([]response.SimpleLabel, len(*p.Labels))
		if p.Labels != nil {
			for i, l := range *p.Labels {
				ls[i] = response.SimpleLabel{
					LabelID: l.ID,
					Content: l.Content,
				}
			}
		}
		rsp[i] = response.SimpleExprPost{
			ID:           p.ID,
			Title:        p.Title,
			Content:      p.Content,
			Column:       p.Column,
			Labels:       &ls,
			CollectorNum: uint(len(*p.Collectors)),
		}
	}
	return &rsp
}
func BuildSimpleDetail(post *table.ExprPost) *response.SimpleExprDetailResp {
	ls := make([]response.SimpleLabel, len(*post.Labels))
	for i, l := range *post.Labels {
		ls[i] = response.SimpleLabel{
			LabelID: l.ID,
			Content: l.Content,
		}
	}
	num1 := 0
	num2 := 0
	if post.Comments != nil {
		num1 = len(*post.Comments)
	}
	if post.Collectors != nil {
		num2 = len(*post.Collectors)
	}
	r := response.SimpleExprDetailResp{
		ID:             post.ID,
		BucuoId:        post.Publisher.BucuoID,
		PublisherTime:  response.Time(post.CreatedAt),
		ProfilePicture: post.Publisher.ProfilePicture,
		Labels:         &ls,
		CollectorNum:   num2,
		CommentNum:     num1,
	}
	return &r
}

func (e ExprPostService) DeleteOne(id uint) error {
	return exprdao.DeleteOne(id)
}
func (e ExprPostService) PushPost(req *request.ExprPostReq) string {
	labels := make([]table.Label, len(req.Labels), len(req.Labels))
	for i, value := range req.Labels {
		labels[i] = table.Label{
			Content: value,
		}
	}
	post := &table.ExprPost{
		Title:       req.Title,
		Content:     req.Content,
		Column:      req.Column,
		PublisherID: req.PublishId,
		Labels:      &labels,
	}
	return exprdao.PostExpr(post)
}
func (e ExprPostService) FindAll(column string, pagesize uint, pagenum uint) (*[]response.SimpleExprPost, error) {
	result, err := exprdao.FindAll(column, pagesize, pagenum)
	if err != nil {
		return nil, err
	}
	return BuildSimpleExprPost(result), nil
}
func (e ExprPostService) FindDetails(id uint) (error, *response.SimpleExprDetailResp) {
	post, err := exprdao.FindOne(id)
	if err != nil {
		return err, nil
	}
	return nil, BuildSimpleDetail(post)
}
func (e ExprPostService) ExprExist(exprid uint, uid uint) error {
	return exprdao.ExistOne(exprid, uid)
}

func (e ExprPostService) UpdateOne(req *request.UpdateExprReq) error {
	return exprdao.UpdateOne(req)
}
