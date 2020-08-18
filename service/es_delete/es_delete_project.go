/**
 * @Description:
 * @File: es_delete
 * @Date: 2020/8/18 0018 10:49
 */

package es_delete

import "csxft/serializer"

//删除楼盘服务
type DeleteProjectService struct {
	ProjectId  int `form:"project_id" json:"project_id" binding:"required"`
}

func (service *DeleteProjectService) DeleteProject() serializer.Response {

	DeleteDoc(service.ProjectId, "project")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

