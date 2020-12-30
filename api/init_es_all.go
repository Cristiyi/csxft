/**
 * @Description:
 * @File: init_es_all
 * @Date: 2020/9/9 0009 10:47
 */

package api

import (
	service "csxft/service/es"
	"github.com/gin-gonic/gin"
)

//初始化所有楼盘数据
func InitAllProject(c *gin.Context) {
	var service service.InitProjectAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitProjectAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有批次数据
func InitAllBatch(c *gin.Context) {
	var service service.InitBatchAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitBatchAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有一房一价数据
func InitAllHouse(c *gin.Context) {
	var service service.InitHouseAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitHouseAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有楼盘动态数据
func InitAllDynamic(c *gin.Context) {
	var service service.InitDynamicAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitDynamicAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有摇号结果数据
func InitAllLotteryResult(c *gin.Context) {
	var service service.InitLotteryResultAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitLotteryAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}
//初始化所有认筹结果数据
func InitAllSolicitResult(c *gin.Context) {
	var service service.InitSolicitResultAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitSolicitResultAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}


//初始化所有公告数据
func InitAllNoticeResult(c *gin.Context) {
	var service service.InitNoticeAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitNoticeAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有公告数据
func InitAllHouseTypeImageResult(c *gin.Context) {
	var service service.InitHouseTypeImageAllService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitHouseTypeImageAll()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有批次-楼盘数据
func InitBatchProject(c *gin.Context) {
	var service service.InitBatchProjectService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitBatchProject()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化所有批次-楼盘数据
func InitBatchProjectByProject(c *gin.Context) {
	var service service.InitBatchProjectByProjectService
	if err := c.ShouldBind(&service); err == nil {
		res := service.InitBatchProjectByProject()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}