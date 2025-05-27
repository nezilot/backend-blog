package route

import (
	"blog/boostrap"
	"blog/pgsql"
	"github.com/gin-gonic/gin"
)

func Setup(env *boostrap.Env, gin *gin.Engine, db pgsql.Database) {
	publicRouter := gin.Group("")

	// public api
	NewCommentRouter(env, db, publicRouter)

	// privates api
	NewPostRouter(publicRouter, db)
}
