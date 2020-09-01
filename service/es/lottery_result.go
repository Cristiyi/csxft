/**
 * @Description:
 * @File: lottery_result
 * @Date: 2020/7/23 0023 15:40
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"reflect"
	"strconv"
)

//搜索摇号结果服务
type LotteryResultService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort_type" json:"sort_type"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	Search    string `form:"search" json:"search"`
	Status int32 `form:"status" json:"status"`
}


func (service *LotteryResultService) GetLotteryResult() serializer.Response {

	commonParam := make(map[string]string)
	batch := GetTargetBatch(service.ProjectId, service.Status)
	commonParam["ProjectId"] = service.ProjectId
	commonParam["BatchId"] = strconv.Itoa(int(batch.ID))
	if service.Sort != "" {
		commonParam["sort"] = service.Sort
	} else {
		commonParam["sort"] = "Type"
	}
	if service.SortType != "" {
		commonParam["sortType"] = service.SortType
	} else {
		commonParam["sortType"] = "desc"
	}

	if service.Search != "" {
		commonParam["Search"] = service.Search
	}

	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 10
	}
	res := QueryLotteryResult(service.Start, size, commonParam)
	if res != nil && len(res.Hits.Hits)>0 {
		var result []model.LotteryResult
		for _, item := range res.Each(reflect.TypeOf(model.LotteryResult{})) {
			if t, ok := item.(model.LotteryResult); ok {
				//t.IdCardBack = ""
				t.TotalHouse = batch.AllNo
				t.TotalPerson = batch.LotteryNo
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
			Code: 400,
			Msg: "暂无数据",
		}
	}

}