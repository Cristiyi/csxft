/**
 * @Description:
 * @File: area
 * @Date: 2020/7/23 0023 19:25
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"reflect"
)

//搜索摇号结果服务
type CsAreaService struct {
}

func (service *CsAreaService) GetCsArea() serializer.Response {
	res := GetCsArea()
	if res != nil && len(res.Hits.Hits)>0 {
		var result []model.Area
		for _, item := range res.Each(reflect.TypeOf(model.Area{})) {
			if t, ok := item.(model.Area); ok {
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
