package router

import (
	"backend/api"
	"github.com/gin-gonic/gin"
)

func ArticleRouter(router *gin.Engine) {
	articleApi := api.Api.ArticleApi
	article := router.Group("article")
	{
		article.POST("/add", articleApi.AddArticle)
	}
}
