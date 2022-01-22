package service

import (
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/response"
	"bucuo/model/table"
	"strings"
)

var commondao = dao.CommonDao{}

type CommonService struct {
}

func (c CommonService) BuildSimpleSkillPosts(rs *[]table.SkillPost) *[]response.SimpleCommonPost {
	res := make([]response.SimpleCommonPost, len(*rs))
	for i, p := range *rs {
		ls := make([]response.SimpleLabel, len(*p.Labels))
		if p.Labels != nil {
			for i, l := range *p.Labels {
				ls[i] = response.SimpleLabel{
					LabelID: l.ID,
					Content: l.Content,
				}
			}
		}
		rspath := "/static/images/default.png"
		if p.Resources != nil {
			for _, resource := range *p.Resources {
				if strings.HasSuffix(resource.DiskFilePath, ".jpg") ||
					strings.HasSuffix(resource.DiskFilePath, ".png") {
					rspath = "/api/resource/" + resource.ID
				}
			}
		}
		res[i] = response.SimpleCommonPost{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			Column:  p.Column,
			Labels:  &ls,
			Cover:   rspath,
		}
	}
	return &res
}
func (c CommonService) BuildSimpleLocalPosts(rs *[]table.LocalPost) *[]response.SimpleCommonPost {
	res := make([]response.SimpleCommonPost, len(*rs))
	for i, p := range *rs {
		ls := make([]response.SimpleLabel, len(*p.Labels))
		if p.Labels != nil {
			for i, l := range *p.Labels {
				ls[i] = response.SimpleLabel{
					LabelID: l.ID,
					Content: l.Content,
				}
			}
		}
		rspath := "/static/images/default.png"
		if p.Resources != nil {
			for _, resource := range *p.Resources {
				if strings.HasSuffix(resource.DiskFilePath, ".jpg") ||
					strings.HasSuffix(resource.DiskFilePath, ".png") {
					rspath = "/api/resource/" + resource.ID
				}
			}
		}
		res[i] = response.SimpleCommonPost{
			ID:      p.ID,
			Title:   p.Title,
			Content: p.Content,
			Column:  p.Column,
			Labels:  &ls,
			Cover:   rspath,
		}
	}
	return &res
}

func (c CommonService) BuildSimpleSkillPostDetail(post *table.SkillPost) *response.SimpleCommonPostDetailResp {
	ls := make([]response.SimpleLabel, len(*post.Labels))
	rs := make([]string, len(*post.Labels))
	for i := 0; i < len(ls); i++ {
		ls[i] = response.SimpleLabel{
			LabelID: (*(post.Labels))[i].ID,
			Content: (*(post.Labels))[i].Content,
		}
		rs[i] = "/api/resource/" + (*(post.Resources))[i].ID
	}
	num1 := 0
	if post.Comments != nil {
		num1 = len(*post.Comments)
	}
	r := response.SimpleCommonPostDetailResp{
		ID:             post.ID,
		BucuoId:        post.Publisher.BucuoID,
		PublisherTime:  response.Time(post.CreatedAt),
		ProfilePicture: post.Publisher.ProfilePicture,
		Labels:         &ls,
		CommentNum:     num1,
		Resources:      &rs,
	}
	return &r

}
func (c CommonService) BuildSimpleLocaLPostDetail(post *table.LocalPost) *response.SimpleCommonPostDetailResp {
	ls := make([]response.SimpleLabel, len(*post.Labels))
	rs := make([]string, len(*post.Resources))
	for i, label := range *post.Labels {
		ls[i] = response.SimpleLabel{
			LabelID: label.ID,
			Content: label.Content,
		}
	}

	for i, resource := range *post.Resources {
		rs[i] = "/api/resource/" + resource.ID
	}
	num1 := 0
	if post.Comments != nil {
		num1 = len(*post.Comments)
	}
	r := response.SimpleCommonPostDetailResp{
		ID:             post.ID,
		BucuoId:        post.Publisher.BucuoID,
		PublisherTime:  response.Time(post.CreatedAt),
		ProfilePicture: post.Publisher.ProfilePicture,
		Labels:         &ls,
		CommentNum:     num1,
		Resources:      &rs,
	}
	return &r

}
func (c CommonService) PushPost(req *request.CommonPostReq) string {

	err, id := commondao.CreatePost(req.Type, req.Title, req.Content, req.Column, req.PublishId, &req.Labels)
	if err != nil {
		return err.Error()
	}
	err = commondao.AddResource(req.Type, id, &req.Resources)
	if err != nil {
		return err.Error()
	}
	return ""
}
func (c CommonService) FindAll(posttype string, column string, pagesize uint, pagenum uint) (string, *[]response.SimpleCommonPost) {
	if posttype == "local_posts" {
		ps, err := commondao.FindAllLocalPost(posttype, column, pagesize, pagenum)
		if err != nil {
			return err.Error(), nil
		}
		return "", c.BuildSimpleLocalPosts(ps)
	}
	if posttype == "skill_posts" {
		ps, err := commondao.FindAllSkillPost(posttype, column, pagesize, pagenum)
		if err != nil {
			return err.Error(), nil
		}
		return "", c.BuildSimpleSkillPosts(ps)
	}
	return "未知错误", nil
}
func (c CommonService) Exist(req *request.DeleteCommonReq, uid uint) bool {
	return commondao.Exist(req.Type, req.PostID, uid)
}
func (c CommonService) Delete(posttype string, id uint, uid uint) string {
	return commondao.Delete(posttype, id, uid)
}
func (c CommonService) FindDetail(req *request.DeleteCommonReq) (*response.SimpleCommonPostDetailResp, string) {
	if req.Type == "skill_posts" {
		r, err := commondao.FindOneSkillPost(req.PostID)
		if err != nil {
			return nil, err.Error()
		}
		return c.BuildSimpleSkillPostDetail(r), ""
	}
	if req.Type == "local_posts" {
		r, err := commondao.FindOneLocalPost(req.PostID)
		if err != nil {
			return nil, err.Error()
		}
		return c.BuildSimpleLocaLPostDetail(r), ""
	}
	return nil, "未知错误"
}
