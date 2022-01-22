package main

import (
	"bucuo/constant/exprpostcolumn"
	"bucuo/dao"
	"bucuo/model/request"
	"bucuo/model/response"
	"bucuo/model/table"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

//dsn = root:yeep2020@tcp(110.42.190.71:3307)/bucuo?charset=utf8mb4&parseTime=True&loc=Local
func TestCreateExprPost(t *testing.T) {
	r := table.ExprPost{
		Title:       "测试2",
		Content:     "测试测试内容2",
		Column:      exprpostcolumn.COMPETITION,
		PublisherID: 1,
		Labels: &[]table.Label{{
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
	r := table.Comment{
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
	r := table.Resource{
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

	r := []table.User{{
		Sno:      "1904060080",
		Password: "yeep2020",
		Username: "yeep1",
		OpenId:   "1",
		College:  "河南大学",
		Phone:    "18712337064",
	},
		{
			Sno:      "1904060081",
			Password: "yeep2020",
			Username: "yeep2",
			OpenId:   "2",
			College:  "河南大学",
			Phone:    "18712337064",
		},
		{
			Sno:      "1904060082",
			Password: "yeep2020",
			Username: "yeep3",
			OpenId:   "3",
			College:  "河南大学",
			Phone:    "18712337064",
		}, {
			Sno:      "1904060083",
			Password: "yeep2020",
			Username: "yeep4",
			OpenId:   "4",
			College:  "河南大学",
			Phone:    "18712337064",
		},
	}
	err := dao.DB.Create(r).Error
	fmt.Printf("%#v", r)
	fmt.Printf("%#v", err)

}
func TestCreateReply(t *testing.T) {
	r := table.Reply{
		Content: "12313113",
	}
	dao.DB.Create(&r)
	fmt.Printf("%#v", r)
}
func TestMap(t *testing.T) {
	m := map[string]int{}
	m["1"] = 22
	print()
}
func TestAddHomeLike(t *testing.T) {
	dao.DB.Model(&table.User{
		Model: table.Model{
			ID: 1,
		},
	}).
		Association("LikeHomePageUsers").
		Append(&table.User{
			Model: table.Model{ID: 2}},
		)
	dao.DB.
		Model(&table.User{
			Model: table.Model{
				ID: 1,
			},
		}).
		Association("LikeHomePageUsers").
		Append(&table.User{Model: table.Model{ID: 3}})
	dao.DB.
		Model(&table.User{
			Model: table.Model{
				ID: 2,
			},
		}).
		Association("LikeHomePageUsers").
		Append(&table.User{Model: table.Model{ID: 1}})
	dao.DB.
		Model(&table.User{
			Model: table.Model{
				ID: 3,
			},
		}).
		Association("LikeHomePageUsers").
		Append(&table.User{Model: table.Model{ID: 1}})
	dao.DB.
		Model(&table.User{
			Model: table.Model{
				ID: 4,
			},
		}).
		Association("LikeHomePageUsers").
		Append(&table.User{Model: table.Model{ID: 1}})
}
func TestGetUser(t *testing.T) {
	user := response.UserGetResp{}
	dao.DB.Model(&table.User{
		Model: table.Model{ID: 1},
	}).
		Select("Sno", "Grade", "College", "BucuoID", "IsExprVip",
			"IsStarVip", "VipExpr", "VipStar", "KnowledgeCurrency").
		Scan(&user)
	user.HomeLikeNum = dao.DB.Model(&table.User{
		Model: table.Model{ID: 1},
	}).
		Association("LikeHomePageUsers").
		Count()
	fmt.Printf("%#v", user)
}
func TestUpdate(t *testing.T) {
	s := dao.DB.Model(&table.User{
		Model: table.Model{
			ID: 1,
		},
	}).Omit("OpenId").Updates(&request.UserCreateReq{Sno: "123"}).Error
	fmt.Println(s)
}
func TestGetHomeLikeNum(t *testing.T) {
	//user := response.UserGetResp{}
	//dao.DB.table.Model(&table.User{
	//	table.Model: table.Model{ID: 2},
	//}).
	//	Select("Sno", "Grade", "College", "BucuoID", "IsExprVip",
	//		"IsStarVip", "VipExpr", "VipStar", "KnowledgeCurrency").
	//	Scan(&user)
	//users := []table.User{}
	//sqlDB.Exec("SELECT COUNT(*) FROM home_page_like_liked where like_home_page_user_id=?")
	var i int64
	r := dao.DB.Raw("SELECT COUNT(*) FROM home_page_like_liked where like_home_page_user_id=?", 2).Row()
	r.Scan(&i)
	fmt.Println(i)
}
func TestCreateExpr(t *testing.T) {
	posts := []table.ExprPost{{
		Title:   "测试",
		Content: "我是一只小小小鸟怎么飞也飞不高！",
		Column:  "课程考试",
		Labels: &[]table.Label{
			{
				Content: "计算机与信息工程学院",
			},
			{
				Content: "数学院",
			},
		},
		PublisherID: 1,
	}, {
		Title:   "测试测试",
		Content: "两只小蜜蜂啊飞到花丛中啊飞啊飞啊！",
		Column:  "课程考试",
		Labels: &[]table.Label{
			{
				Content: "文学院",
			},
			{
				Content: "历史学院",
			},
		},
		PublisherID: 2,
	}}
	dao.DB.Model(&table.ExprPost{}).Create(posts)
}
func TestFindAll(t *testing.T) {
	result := []map[string]interface{}{}
	//exprmap := []request.ExprReq{}
	//exprmap := []table.ExprPost{}
	//exprmap[1].Labels[0] = table.Label{Content: ""}
	////	Table("expr_posts").
	////	Joins("left join labels on labels.owner_id = expr_posts.id").
	////	Find(&exprmap)
	////r := []req{}
	//fmt.Printf("%#v\n", exprmap)
	dao.DB.
		Table("expr_posts").
		//Preload("Labels").
		Find(&result)
	r, _ := json.Marshal(&result)
	fmt.Printf("%#v\n", string(r))
	//
	//bs, _ := json.Marshal()
	//postlist, _ := dao.ExprDao{}.FindAll(1, 1)
	//fmt.Printf("%#v", string(bs))
}
func TestJsonMarsh(t *testing.T) {
	type User struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	type PublicUser struct {
		*User           // 匿名嵌套
		Password string `json:"password,omitempty"`
	}
	u1 := User{
		Name:     "七米",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b) // str:{"name":"七米"}

}
func TestAddLocalPost(t *testing.T) {
	posts := []table.LocalPost{{
		Title:   "测试",
		Content: "我是一只小小小鸟怎么飞也飞不高！",
		Column:  "美食",
		Labels: &[]table.Label{
			{
				Content: "计算机与信息工程学院",
			},
			{
				Content: "数学院",
			},
		},
		PublisherID: 1,
	}, {
		Title:   "测试测试",
		Content: "两只小蜜蜂啊飞到花丛中啊飞啊飞啊！",
		Column:  "美食",
		Labels: &[]table.Label{
			{
				Content: "文学院",
			},
			{
				Content: "历史学院",
			},
		},
		PublisherID: 2,
	}}
	dao.DB.Create(posts)
}
func TestExist(t *testing.T) {
	r := request.ResourceReq{
		OwnerID:   1,
		OwnerType: "local_posts",
	}
	x, _ := dao.ResourceDao{}.PostExist(r)
	fmt.Printf("%#v", x)
}
func TestGetCount(t *testing.T) {
	i, _ := dao.ResourceDao{}.GetPostCount(request.ResourceReq{
		OwnerID:   1,
		OwnerType: "local_posts",
	})
	fmt.Println(i)
}
func TestSlice(t *testing.T) {
	ss := []string{}
	ss = append(ss, "123")
}
func TestUploader(t *testing.T) {
	r := table.Resource{ID: "25d00ed8-7aa8-11ec-9caa-040e3cce8655"}
	dao.DB.
		Table("resources").
		Preload("Uploader").
		First(&r)
	fmt.Printf("%#v", r)
}
func TestExpr(t *testing.T) {
	p := []table.ExprPost{}
	dao.DB.Table("expr_posts").
		Preload("Publisher").
		Preload("Labels").
		Find(&p)
	fmt.Printf("%#v", p[0].Labels)
}
func TestJSon(t *testing.T) {
	r := request.DeleteCommonReq{}
	bs, _ := json.Marshal(r)
	fmt.Println(string(bs))
}
func TestExprOne(t *testing.T) {
	v, _ := dao.ExprDao{}.FindOne(1)
	fmt.Printf("%#v\n", v.Labels)
	fmt.Printf("%#v\n", v.Publisher)
}
func TestDelete(t *testing.T) {
	e := table.ExprPost{Model: table.Model{ID: 2}}
	dao.DB.Unscoped().Delete(&e)
}
func TestFindComonALl(t *testing.T) {
	rs := &[]table.SkillPost{}
	var ro interface{} = rs
	dao.DB.
		Table("skill_posts").
		Limit(int(5)).
		Offset(int((1 - 1) * 2)).
		Preload("Resources").
		Preload("Labels").
		Find(rs)
	fmt.Printf("%#v\n", (*rs)[0])
	fmt.Printf("%#v\n", (*rs)[0].Labels)
	fmt.Printf("%#v\n", (*rs)[0].Resources)
	fmt.Println(ro.(*table.CommonPost))
}
func TestDeletePost(t *testing.T) {
	//v := dao.DB.Exec("delete from "+"skill_posts"+" where id=? and publisher_id=?", "2", "1")
	//fmt.Println(v.Error)
	//fmt.Println(v.RowsAffected)
	dao.DB.Table("local_posts").Where("id=? and publisher_id=?", 12, 1).Delete(table.LocalPost{})
}
