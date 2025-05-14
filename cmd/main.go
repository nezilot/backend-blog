package main

import (
	"blog/router"
	"github.com/gin-gonic/gin"
)

func main() {

	gin := gin.Default()

	route.Setup(gin)

	gin.Run(":3002")

}
