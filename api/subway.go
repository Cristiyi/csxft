/**
 * @Description:
 * @File: subway
 * @Date: 2020/8/7 0007 17:41
 */

package api

import (
	service "csxft/service/es"
	"github.com/gin-gonic/gin"
)

//获取沿线地铁站楼盘数量
func GetSubwayProjectCount(c *gin.Context) {
	var service service.SubwayProjectCountService
	if err := c.ShouldBind(&service); err == nil {
		res := service.GetSubwayProjectCount()
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

