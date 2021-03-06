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

type AreaResult struct{
	ID uint
	Name string
	Pid int64
	ProjectCount int64
}

//长沙区域结果服务
type CsAreaService struct {
}

func (service *CsAreaService) GetCsArea() serializer.Response {
	res := GetCsArea()
	if res != nil && len(res.Hits.Hits)>0 {
		var result []AreaResult
		for _, item := range res.Each(reflect.TypeOf(model.Area{})) {
			if t, ok := item.(model.Area); ok {
				if t.ID != 48941 && t.ID != 48943 {
					projectCount := QueryProjectAreaCount(t.ID)
					tempResult := AreaResult{
						ID:           t.ID,
						Name:         t.Name,
						Pid:          t.Pid,
						ProjectCount: projectCount,
					}
					result = append(result, tempResult)
				}
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
