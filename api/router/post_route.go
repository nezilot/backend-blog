package route

import (
	"blog/api/controller"
	"blog/pgsql"

	"github.com/gin-gonic/gin"
)

func NewPostRouter(group *gin.RouterGroup, db pgsql.Database) {
	pc := &controller.PostController{}

	postGroup := group.Group("/post")
	{
		postGroup.GET("", pc.GetPosts)
		postGroup.GET("/:id", pc.GetPost)
		postGroup.POST("", pc.CreatePost)
		postGroup.PUT("/:id", pc.UpdatePost)
		postGroup.DELETE("/:id", pc.DeletePost)
	}

}
