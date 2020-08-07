/**
 * @Description:
 * @File: subway
 * @Date: 2020/8/7 0007 17:21
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"csxft/util"
	"encoding/json"
	"fmt"
	"reflect"
)

//地铁房数量服务
type SubwayProjectCountService struct {
	LinePoints    string `form:"line_points" json:"line_points" binding:"required"`
}

//地铁房数量服务
type SubwayProjectService struct {
	LinePoint    string `form:"line_point" json:"line_point" binding:"required"`
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

type SubwayProject struct {
	Id string
	ProjectName string
	Longitude float64
	Latitude float64
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

}

func (service *SubwayProjectService) GetSubwayProject() serializer.Response {

	var linePointResult LinePoint
	if err := json.Unmarshal([]byte(service.LinePoint), &linePointResult);err != nil{
		fmt.Println(err)
		return serializer.Response{
			Code: 400,
			Msg:  "数据有误",
		}
	}else {
		fmt.Println(linePointResult)
		pointRange := util.GetDistancePointRange(linePointResult.Latitude, linePointResult.Longitude, 1)
		var subwayProjectList []*SubwayProject
		esRes := GetProjectByPoint(pointRange)
		if esRes != nil {
			for _, item := range esRes.Each(reflect.TypeOf(model.Project{})) {
				if t, ok := item.(model.Project); ok {
					subwayProject := new(SubwayProject)
					subwayProject.Id = string(t.ID)
					if t.PromotionFirstName != "" {
						subwayProject.ProjectName = t.PromotionFirstName
					} else if t.PromotionSecondName != "" {
						subwayProject.ProjectName = t.PromotionSecondName
					} else {
						subwayProject.ProjectName = t.ProjectName
					}
					subwayProject.Longitude = t.Longitude
					subwayProject.Latitude = t.Latitude
					subwayProjectList = append(subwayProjectList, subwayProject)
				}
			}
			return serializer.Response{
				Code: 200,
				Data: subwayProjectList,
				Msg:  "success",
			}
		}
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg:  "success",
		}
	}

}
