/**
 * @Description:
 * @File: house
 * @Date: 2020/11/18 0018 10:35
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"reflect"
)

//获取楼盘所有批次服务
type ProjectBatchService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
}

//获取楼盘所有一房一价
type ProjectHouseService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	BatchId      int `form:"batch_id" json:"batch_id"`
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort" json:"sort"`
	BuildNo string `form:"build_no" json:"build_no"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	Status int32 `form:"status" json:"status"`
}

//获取楼盘所有批次
func (service *ProjectBatchService) GetProjectBatch() serializer.Response {

	res := GetBatch(service.ProjectId,0)
	if res != nil && len(res.Hits.Hits)>0 {
		var result []model.Batch
		for _, item := range res.Each(reflect.TypeOf(model.Batch{})) {
			if t, ok := item.(model.Batch); ok {
				result = append(result, t)
			}
		}
		return serializer.Response{
			Code: 200,
			Data: result,
			Msg: "success",
		}
	} else {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}
}

//获取楼盘下所有的的一房一价
func (service *ProjectHouseService) GetProjectHouse() serializer.Response {

	var credIdResult []int
	batches := GetTargetBatchAll(service.ProjectId, service.BatchId)
	if batches == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}

	for _, batch := range batches {
		if batch.Creds != nil && len(batch.Creds) > 0 {
			for _, item := range batch.Creds {
				credIdResult = append(credIdResult, int(item.ID))
			}
		}
		var houseResult []model.House
		houseParam := make(map[string]string)
		if service.Sort != "" {
			houseParam["sort"] = service.Sort
		} else {
			houseParam["sort"] = "FloorNo"
		}
		if service.SortType != "" {
			houseParam["sortType"] = service.SortType
		} else {
			houseParam["sortType"] = "asc"
		}
		if service.BuildNo != "" {
			houseParam["BuildNo"] = service.BuildNo
		}
		var size int = 0
		if service.Size != 0 {
			size = service.Size
		}  else {
			size = 10
		}
		esHouse := GetCredHouse(service.Start, size, houseParam, credIdResult)
		if esHouse != nil {
			for _, item := range esHouse.Each(reflect.TypeOf(model.House{})) {
				if t, ok := item.(model.House); ok {
					if t.TypeId > 0 && len(t.TypeString) == 0 {
						t.TypeString = model.GetTypeNameById(t.TypeId).Name
					}
					houseResult = append(houseResult, t)
				}
			}
		}

		if houseResult != nil {
			return serializer.Response{
				Code: 200,
				Data: houseResult,
				Msg: "success",
			}
		}
	}

	return serializer.Response{
		Code: 200,
		Data: nil,
		Msg: "暂无数据",
	}
}






