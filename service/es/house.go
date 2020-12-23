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
	CredId int  `form:"cred_id" json:"cred_id"`
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
type OneHouse struct {
	Lou string `json:"lou"`
	Price string `json:"price"`
	Price1 string `json:"price1"`
	Price2 string `json:"price2"`
}

type Detail struct {
	Title string `json:"title"`
	Data []OneHouse `json:"data"`
}

type GenResult struct {
	Title string `json:"title"`
	Data []Detail `json:"data"`
}

type ImageResult struct {
	ProjectName string `json:"project_name"`
	List []GenResult `json:"list"`
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

	var result = new(ImageResult)
	var list []GenResult

	batch := GetTargetBatch(service.ProjectId, service.Status, service.BatchId)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}

	hasCredId := false
	if service.CredId != 0 {
		hasCredId = true
	}

	if batch.Creds != nil && len(batch.Creds) > 0 {
		projectId, _ := strconv.Atoi(service.ProjectId)
		project, _ := repo.NewProjectRepo().GetOne(uint64(projectId))
		if project.PromotionFirstName != "" {
			result.ProjectName = project.PromotionFirstName
		} else {
			result.ProjectName = project.ProjectName
		}
		for _, item := range batch.Creds {
			if hasCredId && item.ID != uint(service.CredId) {
				continue
			}
			cred, err := repo.NewCredRepo().GetToEsData(uint64(item.ID))
			if err == nil {
				var tempData = new(GenResult)
				tempData.Title = cred.BuildingNo
				house, err := repo.NewHouseRepo().GetOneCredHouseData(uint64(item.ID))
				if err == nil {
					houseNoGroup := BuildHouseNo(house)
					if houseNoGroup != nil {
						tempData.Data = houseNoGroup
					}
				}
				list = append(list, *tempData)
			}
			result.List = list
		}
		return serializer.Response{
			Code: 200,
			Data: result,
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
func BuildHouseNo(house []model.House) []Detail{

	var result []Detail
	for _, item := range house {
		//判断户室号长度
		if len(item.HouseNo) <= 4 {
			var number string
			if len(item.HouseNo) == 4 {
				number = item.HouseNo[2:]
			} else {
				number = item.HouseNo[1:]
			}
			if number[0] == '0' {
				number = number[1:]
			}
			resultKey := IsContain(result, number)
			if resultKey == -1 {
				var detail = new(Detail)
				detail.Title = number
				var tempHouse = new(OneHouse)
				tempHouse.Lou = strconv.Itoa(item.FloorNo) + "层"
				if item.HouseAcreage != 0 {
					tempHouse.Price = util.Float2String(item.HouseAcreage, 64)
					tempHouse.Price = tempHouse.Price + "m²"
				}
				if item.UnitPrice != 0 {
					tempHouse.Price1 = util.Float2String(item.UnitPrice, 64)
					tempHouse.Price1 = tempHouse.Price1 + "元/m²"
				}
				if item.HouseAcreage != 0 && item.UnitPrice != 0 {
					tempTotalPrice := item.HouseAcreage * item.UnitPrice / 10000
					tempTotalPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tempTotalPrice), 64)
					tempHouse.Price2 = util.Float2String(tempTotalPrice, 64)
					tempHouse.Price2 = tempHouse.Price2 + "万"
				}
				detail.Data = append(detail.Data, *tempHouse)
				result = append(result, *detail)
			} else {
				var tempHouse = new(OneHouse)
				tempHouse.Lou = strconv.Itoa(item.FloorNo) + "层"
				if item.HouseAcreage != 0 {
					tempHouse.Price = util.Float2String(item.HouseAcreage, 64)
					tempHouse.Price = tempHouse.Price + "m²"
				}
				if item.UnitPrice != 0 {
					tempHouse.Price1 = util.Float2String(item.UnitPrice, 64)
					tempHouse.Price1 = tempHouse.Price1 + "元/m²"
				}
				if item.HouseAcreage != 0 && item.UnitPrice != 0 {
					tempTotalPrice := item.HouseAcreage * item.UnitPrice / 10000
					tempTotalPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", tempTotalPrice), 64)
					tempHouse.Price2 = util.Float2String(tempTotalPrice, 64)
					tempHouse.Price2 = tempHouse.Price2 + "万"
				}
				result[resultKey].Data = append(result[resultKey].Data, *tempHouse)
			}
		}
	}
	return result

}

//判断数组是否存在某值
func IsContain(items []Detail, item string) int {
	for i, eachItem := range items {
		if eachItem.Title == item {
			return i
		}
	}
	return -1
}






