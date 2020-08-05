/**
 * @Description:
 * @File: house_type_image
 * @Date: 2020/7/24 0024 12:13
 */

package es

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"reflect"
	"strconv"
)

//获取户型图服务
type GetHouseTypeImageService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	HomeNum    string `form:"home_num" json:"home_num"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	Status int32 `form:"status" json:"status"`
}

//获取户型图数量服务
type GetHouseTypeImageNumService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Status int32 `form:"status" json:"status"`
}

type HouseTypeImageResult struct {
	Text string
	Num int64
}

func (service *GetHouseTypeImageNumService) GetGroup() serializer.Response {
	batch := GetTargetBatch(service.ProjectId, service.Status)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Msg: "暂无数据",
		}
	}
	projectId, _ := strconv.Atoi(service.ProjectId)
	res := repo.NewHouseTypeImageRepo().GetHouseImageGroup(uint64(projectId), batch.ID)
	var data []HouseTypeImageResult
	if res != nil {
		allTypeName := repo.HouseImageGroup{
			HomeNum:"全部",
		}
		var groupRes []repo.HouseImageGroup
		groupRes = append(groupRes, allTypeName)
		groupRes = append(groupRes, res...)
		commonParam := make(map[string]string)
		commonParam["ProjectId"] = service.ProjectId
		for _, item := range groupRes {
			if item.HomeNum == "全部" {
				commonParam["HomeNum"] = ""
			} else {
				commonParam["HomeNum"] = item.HomeNum
			}

			num := HouseTypeImageCount(commonParam)
			houseTypeImageResult := HouseTypeImageResult{
				Text: item.HomeNum,
				Num: num,
			}
			data = append(data, houseTypeImageResult)

		}
		return serializer.Response{
			Code: 200,
			Data: data,
			Msg: "success",
		}
	}
	return serializer.Response{
		Code: 200,
		Msg: "暂无数据",
	}
}

func (service *GetHouseTypeImageService) GetHouseTypeImage() serializer.Response {
	batch := GetTargetBatch(service.ProjectId, service.Status)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Msg: "暂无数据",
		}
	}
	commonParam := make(map[string]string)
	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 10
	}
	commonParam["ProjectId"] = service.ProjectId
	commonParam["BatchId"] = strconv.Itoa(int(batch.ID))
	if service.HomeNum != "" {
		commonParam["HomeNum"] = service.HomeNum
	}
	commonParam["sort"] = "UpdatedAt"
	commonParam["sortType"] = "desc"
	if service.HomeNum != "" {
		commonParam["HomeNum"] = service.HomeNum
	}
	res := QueryHouseTypeImage(service.Start, size, commonParam)
	if res != nil {
		var result []model.HouseTypeImage
		for _, item := range res.Each(reflect.TypeOf(model.HouseTypeImage{})) {
			if t, ok := item.(model.HouseTypeImage); ok {
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
			Msg: "暂无数据",
		}
	}
}


