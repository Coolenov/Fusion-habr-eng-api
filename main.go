package main

import (
	"Habr-ru-api/api"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/get/news", api.ReturnAllPosts)
	r.Run(":8020")
}
