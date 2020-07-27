/**
 * @Description:
 * @File: dynamic
 * @Date: 2020/7/23 0023 10:14
 */

package es

import (
	"CMD-XuanFangTong-Server/model"
	"CMD-XuanFangTong-Server/serializer"
	"reflect"
)

//获取楼盘动态数量服务
type DynamicCountService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	//Type int32 `form:"type" json:"type"`
}

//获取楼盘动态服务
type DynamicService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Type string `form:"type" json:"type"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort_type" json:"sort_type"`
}

//获取楼盘动态数量
func (service DynamicCountService) GetDynamicCount() serializer.Response {
	countParam := make(map[string]string)
	data := make(map[string]interface{})
	countParam["ProjectId"] = service.ProjectId
	//个数
	countParam["type"] = ""
	data["allCount"] = QueryDynamicCount(countParam)
	countParam["type"] = "1"
	data["newsCount"] = QueryDynamicCount(countParam)
	countParam["type"] = "2"
	data["preSellCount"] = QueryDynamicCount(countParam)
	countParam["type"] = "3"
	data["chooseHouseCount"] = QueryDynamicCount(countParam)

	return serializer.Response{
		Code: 200,
		Data: data,
		Msg: "success",
	}
}


func (service DynamicService) GetDynamic() serializer.Response {
	commonParam := make(map[string]string)
	commonParam["ProjectId"] = service.ProjectId
	if service.Sort != "" {
		commonParam["sort"] = service.Sort
	} else {
		commonParam["sort"] = "CreatedAt"
	}
	if service.SortType != "" {
		commonParam["sortType"] = service.SortType
	} else {
		commonParam["sortType"] = "desc"
	}
	if service.Type != "" {
		commonParam["type"] = service.Type
	}

	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 10
	}
	res := QueryDynamic(service.Start, size, commonParam)
	if res != nil {
		var result []model.Dynamic
		for _, item := range res.Each(reflect.TypeOf(model.Dynamic{})) {
			if t, ok := item.(model.Dynamic); ok {
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



