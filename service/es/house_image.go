/**
 * @Description:
 * @File: house_image
 * @Date: 2020/12/18 0018 15:25
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"reflect"
	"strconv"
)

//获取一房一价图服务
type GetHouseImageService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Status int32 `form:"status" json:"status"`
	BatchId int `form:"batch_id" json:"batch_id"`
}

//获取最近开盘的一房一价图
func (service GetHouseImageService) GetHouseImage() serializer.Response {

	batch := GetTargetBatch(service.ProjectId, service.Status, service.BatchId)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}
	commonParam := make(map[string]string)
	commonParam["BatchId"] = strconv.Itoa(int(batch.ID))
	noticeRes := GetHouseImage(commonParam)
	if noticeRes != nil {
		var result []model.HouseImage
		for _, item := range noticeRes.Each(reflect.TypeOf(model.HouseImage{})) {
			if t, ok := item.(model.HouseImage); ok {
				result = append(result, t)
			}
		}
		return serializer.Response{
			Code: 200,
			Data: result,
			Msg: "success",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg: "暂无数据",
	}

}
