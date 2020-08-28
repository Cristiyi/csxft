/**
 * @Description:
 * @File: delete_es
 * @Date: 2020/8/18 0018 11:00
 */

package api

import (
	service "csxft/service/es_delete"
	"github.com/gin-gonic/gin"
)

//删除楼盘
func DeleteProject(c *gin.Context) {
	var service service.DeleteProjectService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteProject()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//删除批次
func DeleteBatch(c *gin.Context) {
	var service service.DeleteBatchService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteBatch()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//删除公告
func DeleteNotice(c *gin.Context) {
	var service service.DeleteNoticeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteNotice()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据开盘删除一房一价
func DeleteCredHouse(c *gin.Context) {
	var service service.DeleteCredHouseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteCredHouseService()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据开盘删除一房一价
func DeleteHouse(c *gin.Context) {
	var service service.DeleteHouseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteHouse()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据摇号结果
func DeleteLotteryResult(c *gin.Context) {
	var service service.DeleteLotteryResultService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteLotteryResult()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//根据批次删除摇号结果
func DeleteBatchLotteryResult(c *gin.Context) {
	var service service.DeleteBatchLotteryResultService
	if err := c.ShouldBind(&service); err == nil {
		res := service.DeleteBatchLotteryService()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}