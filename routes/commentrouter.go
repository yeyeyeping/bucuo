package routes

import (
	"bucuo/controller"
	"github.com/gin-gonic/gin"
)

var commentcontroller = controller.CommentController{}

func AuthCommentInit(group *gin.RouterGroup) {
	group.POST("comment", commentcontroller.AddComment)
	group.DELETE("comment/:id", commentcontroller.DeleteComment)
	group.POST("reply", commentcontroller.AddReply)
	group.DELETE("reply/:id", commentcontroller.DeleteReply)
}
func CommentInit(group *gin.RouterGroup) {

}
