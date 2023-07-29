package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jun-chiang/go-web-demo1/service"
)

func QueryPageInfo(c *gin.Context) {
	// 获取URL链接中的ID
	topicIdStr := c.Param("topicId")
	data, err := service.QueryPageInfo(topicIdStr)
	if err != nil {
		c.JSON(http.StatusOK, ServerFailed(err.Error()))
		return
	}
	c.JSON(http.StatusOK, ServerSuccess(data))
}
