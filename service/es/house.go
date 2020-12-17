/**
 * @Description:
 * @File: house
 * @Date: 2020/11/18 0018 10:35
 */

package es

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"fmt"
	"reflect"
	"strconv"
	"csxft/util"
)

//获取楼盘所有批次服务
type ProjectBatchService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
}

//生成一房一价图
type GenHouseImageService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Status int32 `form:"status" json:"status"`
	BatchId int `form:"batch_id" json:"batch_id"`
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

//一房一价图所用数据
type oneHouse struct {
	lou string
	price string
	price1 string
	price2 string
}

type detail struct {
	title string
	data []oneHouse
}

type result struct {
	title string
	data []detail
}

//获取楼盘所有批次
func (service *ProjectBatchService) GetProjectBatch() serializer.Response {

	res := GetBatchAll(service.ProjectId,0)
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

//生成一房一价图
func (service GenHouseImageService) GenHouseImage() serializer.Response {

	var data []result

	batch := GetTargetBatch(service.ProjectId, service.Status, service.BatchId)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}

	if batch.Creds != nil && len(batch.Creds) > 0 {
		for _, item := range batch.Creds {
			cred, err := repo.NewCredRepo().GetToEsData(uint64(item.ID))
			if err == nil {
				var tempData = new(result)
				tempData.title = cred.BuildingNo
				house, err := repo.NewHouseRepo().GetOneCredHouseData(uint64(item.ID))
				if err == nil {
					houseNoGroup := BuildHouseNo(house)
					if houseNoGroup != nil {
						tempData.data = houseNoGroup
					}
				}
				data = append(data, *tempData)
			}
		}
		fmt.Println(data)
		return serializer.Response{
			Code: 200,
			Data: data,
			Msg: "success",
		}
	}
	return serializer.Response{
		Code: 200,
		Data: nil,
		Msg: "暂无数据",
	}
}

//获取独立房号 拼装数据
func BuildHouseNo(house []model.House) []detail{

	var result []detail
	for _, item := range house {
		//判断户室号长度
		if len(item.HouseNo) <= 4 {
			var number string
			if len(item.HouseNo) == 4 {
				number = item.HouseNo[2:]
			} else {
				number = item.HouseNo[1:]
			}
			resultKey := IsContain(result, number)
			if resultKey == -1 {
				var detail = new(detail)
				detail.title = number
				var tempHouse = new(oneHouse)
				tempHouse.lou = strconv.Itoa(item.FloorNo)
				if item.HouseAcreage != 0 {
					tempHouse.price = util.Float2String(item.HouseAcreage, 64)
				}
				if item.HouseAcreage != 0 {
					tempHouse.price = util.Float2String(item.HouseAcreage, 64)
					tempHouse.price = tempHouse.price + "m"
				}
				if item.UnitPrice != 0 {
					tempHouse.price1 = util.Float2String(item.UnitPrice, 64)
					tempHouse.price1 = tempHouse.price1 + "元/m"
				}
				if item.TotalPrice != 0 {
					tempHouse.price2 = util.Float2String(item.TotalPrice, 64)
					tempHouse.price2 = tempHouse.price1 + "万"
				}
				detail.data = append(detail.data, *tempHouse)
				result = append(result, *detail)
			} else {
				var tempHouse = new(oneHouse)
				tempHouse.lou = strconv.Itoa(item.FloorNo)
				if item.HouseAcreage != 0 {
					tempHouse.price = util.Float2String(item.HouseAcreage, 64)
				}
				if item.HouseAcreage != 0 {
					tempHouse.price = util.Float2String(item.HouseAcreage, 64)
					tempHouse.price = tempHouse.price + "m"
				}
				if item.UnitPrice != 0 {
					tempHouse.price1 = util.Float2String(item.UnitPrice, 64)
					tempHouse.price1 = tempHouse.price1 + "元/m"
				}
				if item.TotalPrice != 0 {
					tempHouse.price2 = util.Float2String(item.TotalPrice, 64)
					tempHouse.price2 = tempHouse.price1 + "万"
				}
				result[resultKey].data = append(result[resultKey].data, *tempHouse)
			}
		}
	}
	return result

}

//判断数组是否存在某值
func IsContain(items []detail, item string) int {
	for i, eachItem := range items {
		if eachItem.title == item {
			return i
		}
	}
	return -1
}






