package main

import (
	"bucuo/constant/exprpostcolumn"
	"bucuo/dao"
	"bucuo/model"
	"fmt"
	"log"
	"testing"
)

func TestCreateExprPost(t *testing.T) {
	r := model.ExprPost{
		Title:   "测试2",
		Content: "测试测试内容2",
		Column:  exprpostcolumn.COMPETITION,
		OwnerID: 1,
		Labels: []model.Label{{
			Content: "商学院",
		},
			{
				Content: "测试",
			}},
	}
	dao.DB.Create(&r)
	log.Printf("%#v", r)
}
func TestCreateComment(t *testing.T) {
	r := model.Comment{
		Content:   "真不错",
		UserID:    2,
		OwnerType: "expr_posts",
		OwnerID:   1,
	}
	err := dao.DB.Create(&r).Error
	if err != nil {
		log.Fatalf("错误：%#v", err)
	}
	log.Printf("%#v", r)

}
func TestCreateResource(t *testing.T) {
	r := model.Resource{
		DiskFilePath: "1231",
		UploaderID:   1,
	}
	err := dao.DB.Create(&r).Error
	if err != nil {
		log.Fatalf("错误：%#v", err)
	}
	log.Printf("%#v", r)
	log.Printf("%s", r.ID)
}
func TestCreateUser(t *testing.T) {
	r := []model.User{{
		Sno:      "1904060087",
		Password: "yeep2020",
		Username: "yeep",
		OpenId:   "123",
		College:  "河南大学",
		Phone:    "18712337064",
	},
		{
			Sno:      "1904060011",
			Password: "yeep2020",
			Username: "yeep",
			OpenId:   "1234",
			College:  "河南大学",
			Phone:    "18712337064",
		},
	}
	err := dao.DB.Create(r).Error
	fmt.Printf("%#v", r)
	fmt.Printf("%#v", err)

}
func TestCreateReply(t *testing.T) {
	r := model.Reply{
		Content: "12313113",
	}
	dao.DB.Create(&r)
	fmt.Printf("%#v", r)
}
