package api

import (
	"Habr-ru-api/api/internal"
	"github.com/Coolenov/Fusion-library/types"
	"github.com/gin-gonic/gin"
)

func ReturnAllPosts(c *gin.Context) {

	var posts []types.Post

	posts = internal.HabrScraper()
	c.JSON(200, posts)
}
