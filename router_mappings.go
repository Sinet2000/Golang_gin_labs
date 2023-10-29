package main

import (
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func CreateUrlMappings() {
	Router = gin.Default()
	// Router.Use(controllers.Cors())

	Router.LoadHTMLGlob("templates/*")

	api := Router.Group("/api")
	{
		api.GET("/", showIndexPage)
		api.GET("/article/view/:article_id", getArticle)
	}
}
