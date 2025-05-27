package route

import (
	"blog/api/controller"
	"blog/boostrap"
	"blog/pgsql"
	"github.com/gin-gonic/gin"
)

func NewCommentRouter(env *boostrap.Env, db pgsql.Database, group *gin.RouterGroup) {

	cc := &controller.CommentController{}

	commentgroup := group.Group("/comment")
	{
		commentgroup.GET("/:id", cc.GetComment)
		commentgroup.POST("", cc.CreateComment)
		commentgroup.PUT("/:id", cc.UpdateComment)
		commentgroup.DELETE("/:id", cc.DeleteComment)

	}

}
