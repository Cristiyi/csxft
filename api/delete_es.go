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
