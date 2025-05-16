package main

import (
	"blog/api/router"
	"blog/boostrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := boostrap.App()

	env := app.Env

	gin := gin.Default()

	route.Setup(env, gin)

	gin.Run(env.Port)

}
