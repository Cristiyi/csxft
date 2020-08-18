/**
 * @Description:
 * @File: task
 * @Date: 2020/8/17 0017 14:44
 */

package api

import (
	service "csxft/service/task"
	"github.com/gin-gonic/gin"
)

//最新取证定时处理
func NewCredTask(c *gin.Context) {
	var service service.NewCredTaskService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNewCredTask()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//最新取证到期定时处理
func NotNewCredTask(c *gin.Context) {
	var service service.NotNewCredTaskService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetNotNewCredTask()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//正在认筹定时处理
func NewRecognitionTask(c *gin.Context) {
	var service service.RecognitionService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetRecognitionTask()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
