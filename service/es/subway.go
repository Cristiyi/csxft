/**
 * @Description:
 * @File: subway
 * @Date: 2020/8/7 0007 17:21
 */

package es

import (
	"csxft/serializer"
	"csxft/util"
	"encoding/json"
	"fmt"
)

//搜索摇号结果服务
type SubwayProjectCountService struct {
	LinePoints    string `form:"line_points" json:"line_points" binding:"required"`
}

type LinePoint struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	Title string `json:"title"`
}

type LinePointReturn struct {
	Title string `json:"title"`
	Count int64 `json:"count"`
}


func (service *SubwayProjectCountService) GetSubwayProjectCount() serializer.Response {

	var linePointResult []LinePoint
	if err := json.Unmarshal([]byte(service.LinePoints), &linePointResult);err != nil{
		fmt.Println(err)
		return serializer.Response{
			Code: 400,
			Msg:  "数据有误",
		}
	}else {
		var linePointReturn []*LinePointReturn
		for _, item := range linePointResult {
			pointRange := util.GetDistancePointRange(item.Latitude, item.Longitude, 1)
			count := GetProjectCountByPoint(pointRange)
			temp := new(LinePointReturn)
			temp.Title = item.Title
			temp.Count = count
			linePointReturn = append(linePointReturn, temp)
		}
		return serializer.Response{
			Code: 200,
			Data: linePointReturn,
			Msg:  "success",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg:  "暂无数据",
	}

}
