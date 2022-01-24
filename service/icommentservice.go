package service

import (
	"bucuo/model/request"
	"bucuo/model/response"
)

type ICommentService interface {
	AddComment(req *request.CommentReq, uid uint) string
	DeleteCommemnt(id uint, uid uint) string
	AddReply(req *request.AddReplyReq, uid uint) string
	DeleteReply(u uint, uid uint) string
	Like(uid uint, commentid uint) string
	LikeReply(uid uint, commentid uint) string
	GetComments(comment *request.ByPageComment) (string, *[]response.SimpleComment)
	GetReplies(reply *request.ByPageReply) (string, *[]response.SimpleReply)
}
