/**
 * @Description:
 * @File: batch_project
 * @Date: 2020/12/28 0028 16:13
 */

package es

import (
	"csxft/model"
	"csxft/serializer"
	"reflect"
)

//根据批次搜索楼盘服务
type SearchBatchProjectService struct {
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort" json:"sort"`
	//Status    string `form:"status" json:"status"`
	SaleStatus    string `form:"sale_status" json:"sale_status"`
	Name    string `form:"name" json:"name"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	IsWillCred string `form:"is_will_cred" json:"is_will_cred"`
	IsNewCred string `form:"is_new_cred" json:"is_new_cred"`
	IsRecognition string `form:"is_recognition" json:"is_recognition"`
	IsIottery string `form:"is_iottery" json:"is_iottery"`
	IsSell string `form:"is_sell" json:"is_sell"`
	AreaId string `form:"area_id" json:"area_id"`
	Renovation string `form:"renovation" json:"renovation"`
	MaxTotalPrice float64 `form:"max_total_price" json:"max_total_price"`
	MinTotalPrice float64 `form:"min_total_price" json:"min_total_price"`
	//TotalPrice float64 `form:"total_price" json:"total_price"`
	MaxPrice float64 `form:"max_price" json:"max_price"`
	MinPrice float64 `form:"min_price" json:"min_price"`
	//Price float64 `form:"price" json:"price"`
	MaxAcreage float64 `form:"max_acreage" json:"max_acreage"`
	MinAcreage float64 `form:"min_acreage" json:"min_acreage"`
	IsDecoration int `form:"renovation" json:"renovation"`
	IsNotDecoration int `form:"is_not_decora tion" json:"is_not_decoration"`
	PredictCredDate int64 `form:"predict_cred_date" json:"predict_cred_date"`
	HasAerialUpload string `form:"has_aerial_upload" json:"has_aerial_upload"`
	IsNearLineOne  int `form:"is_near_line_one" json:"is_near_line_one"`
	IsNearLineTwo  int `form:"is_near_line_two" json:"is_near_line_two"`
	IsNearLineThird int `form:"is_near_line_third" json:"is_near_line_third"`
	IsNearLineFouth  int `form:"is_near_line_fouth" json:"is_near_line_fouth"`
	IsNearLineFifth  int `form:"is_near_line_fifth" json:"is_near_line_fifth"`
	IsNearLineSixth  int `form:"is_near_line_sixth" json:"is_near_line_sixth"`
	IsNearLineSixthTwo  int `form:"is_near_line_sixth_two" json:"is_near_line_sixth_two"`
}

//根据批次搜索楼盘
func (service *SearchBatchProjectService) SearchBatchProjectService() serializer.Response {

	commonParam := make(map[string]string)
	if service.Sort != "" {
		commonParam["sort"] = service.Sort
	} else {
		commonParam["sort"] = "Project.ViewCount"
	}
	if service.SortType != "" {
		commonParam["sortType"] = service.SortType
	} else {
		commonParam["sortType"] = "desc"
	}
	if service.Name != "" {
		commonParam["name"] = service.Name
	}
	if service.IsWillCred != "" {
		commonParam["IsWillCred"] = service.IsWillCred
	}
	if service.IsNewCred != "" {
		commonParam["IsNewCred"] = service.IsNewCred
	}
	if service.IsRecognition != "" {
		commonParam["IsRecognition"] = service.IsRecognition
	}
	if service.IsIottery != "" {
		commonParam["IsIottery"] = service.IsIottery
	}
	if service.AreaId != "" {
		commonParam["AreaId"] = service.AreaId
	}
	if service.Renovation != "" {
		commonParam["Renovation"] = service.Renovation
	}

	if service.HasAerialUpload != "" {
		commonParam["HasAerialUpload"] = service.HasAerialUpload
	}

	calParams := make(map[string]float64)
	//if service.MaxAcreage != 0 {
	//	calParams["MaxAcreage"] = service.MaxAcreage
	//}
	calParams["MaxAcreage"] = service.MaxAcreage
	//if service.MinAcreage != 0 {
	//	calParams["MinAcreage"] = service.MinAcreage
	//}
	calParams["MinAcreage"] = service.MinAcreage
	//if service.MaxTotalPrice != 0 {
	//	calParams["MaxTotalPrice"] = service.MaxTotalPrice
	//}
	calParams["MaxTotalPrice"] = service.MaxTotalPrice
	//if service.MinTotalPrice != 0 {
	//	calParams["MinTotalPrice"] = service.MinTotalPrice
	//}
	calParams["MinTotalPrice"] = service.MinTotalPrice
	//if service.MaxPrice != 0 {
	//	calParams["MaxPrice"] = service.MaxPrice
	//}
	calParams["MaxPrice"] = service.MaxPrice
	//if service.MinPrice != 0 {
	//	calParams["MinPrice"] = service.MinPrice
	//}
	calParams["MinPrice"] = service.MinPrice

	if service.PredictCredDate != 0 {
		calParams["PredictCredDate"] = float64(service.PredictCredDate)
	}
	if service.IsNearLineOne != 0 {
		calParams["IsNearLineOne"] = float64(service.IsNearLineOne)
	}
	if service.IsNearLineTwo != 0 {
		calParams["IsNearLineTwo"] = float64(service.IsNearLineTwo)
	}
	if service.IsNearLineThird != 0 {
		calParams["IsNearLineThird"] = float64(service.IsNearLineThird)
	}
	if service.IsNearLineFouth != 0 {
		calParams["IsNearLineFouth"] = float64(service.IsNearLineFouth)
	}
	if service.IsNearLineFifth != 0 {
		calParams["IsNearLineFifth"] = float64(service.IsNearLineFifth)
	}
	if service.IsNearLineSixth != 0 {
		calParams["IsNearLineSixth"] = float64(service.IsNearLineSixth)
	}
	if service.IsNearLineSixthTwo != 0 {
		calParams["IsNearLineSixthTwo"] = float64(service.IsNearLineSixthTwo)
	}
	calParams["needed"] = 1

	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 200
	}
	res := QueryBatchProject(service.Start, size, commonParam, calParams)
	if res != nil {
		var result []model.Project
		for _, item := range res.Each(reflect.TypeOf(model.Batch{})) {
			if t, ok := item.(model.Batch); ok {
				if IsContainProject(result, t.Project.ID) == -1 && t.Project.ID != 0 {
					result = append(result, t.Project)
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

//判断数组是否存在某值
func IsContainProject(items []model.Project, id uint) int {
	for i, eachItem := range items {
		if eachItem.ID == id {
			return i
		}
	}
	return -1
}

