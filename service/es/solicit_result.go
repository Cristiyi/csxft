/**
 * @Description:
 * @File: solicit_result
 * @Date: 2020/7/23 0023 18:14
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"reflect"
	"strconv"
)

//搜索认筹结果服务
type SolicitResultService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort_type" json:"sort_type"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	Search    string `form:"search" json:"search"`
	Status int32 `form:"status" json:"status"`
	BatchId int `form:"batch_id" json:"batch_id"`
	Type int `form:"type" json:"type"`
}


func (service *SolicitResultService) GetSolicitResult() serializer.Response {

	commonParam := make(map[string]string)
	batch := GetTargetBatch(service.ProjectId, service.Status, service.BatchId)
	commonParam["ProjectId"] = service.ProjectId
	commonParam["BatchId"] = strconv.Itoa(int(batch.ID))
	if service.Sort != "" {
		commonParam["sort"] = service.Sort
	} else {
		commonParam["sort"] = "UpdatedAt"
	}
	if service.SortType != "" {
		commonParam["sortType"] = service.SortType
	} else {
		commonParam["sortType"] = "asc"
	}

	if service.Search != "" {
		commonParam["Search"] = service.Search
	}
	if service.Type != 0 {
		commonParam["Type"] = strconv.Itoa(service.Type)
	}


	var size int = 0
	if service.Size != 0 {
		size = service.Size
	} else {
		size = 10
	}
	res := QuerySolicitResult(service.Start, size, commonParam)
	if res != nil && len(res.Hits.Hits) > 0 {
		var result []model.SolicitResult
		for _, item := range res.Each(reflect.TypeOf(model.SolicitResult{})) {
			if t, ok := item.(model.SolicitResult); ok {
				t.IdCardBack = ""
				result = append(result, t)
			}
		}
		return serializer.Response{
			Code: 200,
			Data: result,
			Msg:  "success",
		}
	} else {
		return serializer.Response{
			Code: 400,
			Msg:  "暂无数据",
		}
	}

}
