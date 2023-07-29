package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jun-chiang/go-web-demo1/controller"
)

func initRouter(r *gin.Engine) {
	apiRouter := r.Group("demo1")

	apiRouter.GET("/queryPageInfo/:topicId", controller.QueryPageInfo)
}
