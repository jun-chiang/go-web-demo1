package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jun-chiang/go-web-demo1/repository"
)

func main() {
	err := repository.InitTopicIndexMap()
	if err != nil {
		fmt.Println("数据库初始化出错：", err.Error())
	}

	r := gin.Default()

	initRouter(r)

	r.Run(":8081")
}
