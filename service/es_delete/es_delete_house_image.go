/**
 * @Description:
 * @File: es_delete_house_image
 * @Date: 2020/12/18 0018 19:34
 */

package es_delete

import "csxft/serializer"

//删除一房一价图
type DeleteHouseImageService struct {
	BatchId  int `form:"batch_id" json:"batch_id" binding:"required"`
}

//删除一房一价图
func (service *DeleteHouseImageService) DeleteService() serializer.Response {

	DeleteDoc(service.BatchId, "house_image")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}
