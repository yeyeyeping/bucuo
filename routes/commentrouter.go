package routes

import (
	"bucuo/controller"
	"github.com/gin-gonic/gin"
)

var commentcontroller = controller.CommentController{}

func AuthCommentInit(group *gin.RouterGroup) {
	group.POST("comment", commentcontroller.AddComment)
	group.DELETE("comment", commentcontroller.DeleteComment)
	group.POST("reply", commentcontroller.AddReply)
	group.DELETE("reply", commentcontroller.DeleteReply)
}
func CommentInit(group *gin.RouterGroup) {

}
