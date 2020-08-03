package api

import (
	service "csxft/service/init"
	"github.com/gin-gonic/gin"
)



//初始化基础数据
func InitBase(c *gin.Context) {
	var service service.InitBaseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化预售证号数据
func InitCred(c *gin.Context) {
	var service service.InitCredService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化开盘房屋数据
func InitHouse(c *gin.Context) {
	var service service.InitHouseService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

//初始化开盘房屋数据
func InitFdc(c *gin.Context) {
	var service service.InitFdcService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Init()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

