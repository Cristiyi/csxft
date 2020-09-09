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
