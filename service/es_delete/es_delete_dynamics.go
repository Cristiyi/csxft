/**
 * @Description:
 * @File: es_delete_dynamics
 * @Date: 2020/9/7 0007 17:41
 */

package es_delete

import "csxft/serializer"

//删除楼盘动态服务
type DeleteDynamicsService struct {
	DynamicsId  int `form:"dynamics_id" json:"dynamics_id" binding:"required"`
}

//删除摇号结果
func (service *DeleteDynamicsService) DeleteDynamicsService() serializer.Response {

	DeleteDoc(service.DynamicsId, "dynamic")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

