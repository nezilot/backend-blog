package route

import (
	"blog/api/controller"

	"github.com/gin-gonic/gin"
)

func NewPostRouter(group *gin.RouterGroup) {
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
