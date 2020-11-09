/**
 * @Description:
 * @File: es_delete_house_type_image
 * @Date: 2020/11/9 0009 17:52
 */

package es_delete

import "csxft/serializer"

//删除户型图服务
type DeleteHouseTypeImageService struct {
	ProjectId  int `form:"project_id" json:"project_id" binding:"required"`
}


//根据楼盘删除户型图服务
func (service *DeleteHouseTypeImageService) DeleteHouseTypeImageService() serializer.Response {

	DeleteHouseTypeImage(service.ProjectId)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

