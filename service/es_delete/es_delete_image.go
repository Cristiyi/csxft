/**
 * @Description:
 * @File: es_delete_image
 * @Date: 2020/12/18 0018 19:59
 */

package es_delete

import "csxft/serializer"

//删除图片
type DeleteImageService struct {
	ID  int `form:"id" json:"id" binding:"required"`
}

//删除图片
func (service *DeleteImageService) DeleteService() serializer.Response {

	DeleteDoc(service.ID, "house_image")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}
