/**
 * @Description:
 * @File: es_delete_house
 * @Date: 2020/8/28 0028 15:42
 */

package es_delete

import "csxft/serializer"

//删除一房一价服务
type DeleteHouseService struct {
	HouseId  int `form:"house_id" json:"house_id" binding:"required"`
}

//删除一房一价服务
type DeleteCredHouseService struct {
	CredId  int `form:"cred_id" json:"cred_id" binding:"required"`
}

//删除一房一价
func (service *DeleteHouseService) DeleteHouse() serializer.Response {

	DeleteDoc(service.HouseId, "house")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

//根据开盘删除一房一价
func (service *DeleteCredHouseService) DeleteHouseService() serializer.Response {

	DeleteCredHouse(service.CredId)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

//根据开盘删除一房一价
func (service *DeleteCredHouseService) DeleteCredHouseService() serializer.Response {

	DeleteCredHouse(service.CredId)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}