package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostController struct {
}

func (pc *PostController) GetPosts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetPosts"})
}
