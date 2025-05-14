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
	}

}
