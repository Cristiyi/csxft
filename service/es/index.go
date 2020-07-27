/**
 * @Description: 首页服务
 * @File: index
 * @Date: 2020/7/10 0010 19:17
 */

package es

import (
	"CMD-XuanFangTong-Server/model"
	"CMD-XuanFangTong-Server/serializer"
	"reflect"
)

//首页数据服务
type IndexService struct {
}

//首页楼盘
func (service *IndexService) Index() serializer.Response {

	countParam := make(map[string]string)
	data := make(map[string]interface{})
	//个数
	countParam["type"] = "1"
	data["willCredCount"] = QueryProjectCount(countParam)
	countParam["type"] = "2"
	data["newCredCount"] = QueryProjectCount(countParam)
	countParam["type"] = "2"
	data["newCredCount"] = QueryProjectCount(countParam)
	countParam["type"] = "3"
	data["recognitionCount"] = QueryProjectCount(countParam)
	countParam["type"] = "4"
	data["newIotteryCount"] = QueryProjectCount(countParam)
	countParam["type"] = "5"
	data["newSellCount"] = QueryProjectCount(countParam)

	willCredListParam := make(map[string]string)
	willCredListParam["sortType"] = "desc"
	willCredListParam["sort"] = "ViewCount"
	//即将取证列表
	willCredListParam["IsWillCred"] = "1"
	var willCredList []model.Project
	willCredRes := QueryProject(0, 3, willCredListParam)
	if willCredRes != nil {
		for _, item := range willCredRes.Each(reflect.TypeOf(model.Project{})) {
			if t, ok := item.(model.Project); ok {
				willCredList = append(willCredList, t)
			}
		}
	}
	data["willCredList"] = willCredList

	//正在认筹列表
	recognitionListParam := make(map[string]string)
	recognitionListParam["sortType"] = "desc"
	recognitionListParam["sort"] = "ViewCount"
	recognitionListParam["IsRecognition"] = "1"
	var recognitionList []model.Project
	recognitionRes:= QueryProject(0, 3, recognitionListParam)
	if recognitionRes != nil {
		for _, item := range recognitionRes.Each(reflect.TypeOf(model.Project{})) {
			if t, ok := item.(model.Project); ok {
				recognitionList = append(recognitionList, t)
			}
		}
	}
	data["recognitionList"] = recognitionList

	//最近摇号列表
	newIotteryListParam := make(map[string]string)
	newIotteryListParam["sortType"] = "desc"
	newIotteryListParam["sort"] = "ViewCount"
	newIotteryListParam["IsIottery"] = "1"
	var newIotteryList []model.Project
	newIotteryRes:= QueryProject(0, 3, newIotteryListParam)
	if newIotteryRes != nil {
		for _, item := range newIotteryRes.Each(reflect.TypeOf(model.Project{})) {
			if t, ok := item.(model.Project); ok {
				newIotteryList = append(newIotteryList, t)
			}
		}
	}
	data["newIotteryList"] = newIotteryList

	//在售楼盘列表
	newSellListParam := make(map[string]string)
	newSellListParam["sortType"] = "desc"
	newSellListParam["sort"] = "ViewCount"
	newSellListParam["IsSell"] = "1"
	var newSellList []model.Project
	newSellRes:= QueryProject(0, 3, newSellListParam)
	if newSellRes != nil {
		for _, item := range newSellRes.Each(reflect.TypeOf(model.Project{})) {
			if t, ok := item.(model.Project); ok {
				newSellList = append(newSellList, t)
			}
		}
	}
	data["newSellList"] = newSellList

	return serializer.Response{
		Code: 200,
		Data: data,
		Msg: "success",
	}
}