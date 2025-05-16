package main

import (
	"blog/api/router"
	"blog/boostrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := boostrap.App()

	env := app.Env

	db := app.Psql.Database(env.DBName)
	defer app.CloseBDConnection()

	gin := gin.Default()

	route.Setup(env, gin, db)

	gin.Run(env.Port)

}
