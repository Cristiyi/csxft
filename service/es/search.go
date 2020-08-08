/**
 * @Description: 搜索
 * @File: search_strategy
 * @Date: 2020/7/9 0009 19:37
 */

package es

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"reflect"
)

//搜索楼盘服务
type SearchProjectService struct {
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort" json:"sort"`
	//Status    string `form:"status" json:"status"`
	Name    string `form:"name" json:"name"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	IsWillCred string `form:"is_will_cred" json:"is_will_cred"`
	IsNewCred string `form:"is_new_cred" json:"is_new_cred"`
	IsRecognition string `form:"is_recognition" json:"is_recognition"`
	IsIottery string `form:"is_iottery" json:"is_iottery"`
	IsSell string `form:"is_sell" json:"is_sell"`
	AreaId string `form:"area_id" json:"area_id"`
	MaxTotalPrice float64 `form:"max_total_price" json:"max_total_price"`
	MinTotalPrice float64 `form:"min_total_price" json:"min_total_price"`
	MaxUnitPrice float64 `form:"max_unit_price" json:"max_unit_price"`
	MinUnitPrice float64 `form:"min_unit_price" json:"min_unit_price"`
	MaxAcreage float64 `form:"max_acreage" json:"max_acreage"`
	MinAcreage float64 `form:"min_acreage" json:"min_acreage"`
	IsDecoration int `form:"is_decoration" json:"is_decoration"`
	IsNotDecoration int `form:"is_not_decoration" json:"is_not_decoration"`
	IsNearLineOne  int `form:"is_near_line_one" json:"is_near_line_one"`
	IsNearLineTwo  int `form:"is_near_line_two" json:"is_near_line_two"`
	IsNearLineThird int `form:"is_near_line_third" json:"is_near_line_third"`
	IsNearLineFouth  int `form:"is_near_line_fouth" json:"is_near_line_fouth"`
	IsNearLineFifth  int `form:"is_near_line_fifth" json:"is_near_line_fifth"`
	IsNearLineSixth  int `form:"is_near_line_sixth" json:"is_near_line_sixth"`
	PredictCredDate int64 `form:"predict_cred_date" json:"predict_cred_date"`
}

//获取楼盘服务
type GetProjectService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
}

//根据楼盘获取开盘服务
type GetProjectCredService struct {
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort_type" json:"sort_type"`
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
}

//根据开盘获取房屋服务
type GetCredHouseService struct {
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort_type" json:"sort_type"`
	CredId    string `form:"cred_id" json:"cred_id" binding:"required"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
}

//即将取证时间分组服务
type GetPredictCredDate struct {
}

//热门搜索
type GetHotService struct {
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
}

type HotProject struct {
	ProjectName string
	ProjectId uint
}

//搜索楼盘
func (service *SearchProjectService) SearchProject() serializer.Response {

	commonParam := make(map[string]string)
	if service.Sort != "" {
		commonParam["sort"] = service.Sort
	} else {
		commonParam["sort"] = "ViewCount"
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
	if service.IsSell != "" {
		commonParam["IsSell"] = service.IsSell
	}
	if service.AreaId != "" {
		commonParam["AreaId"] = service.AreaId
	}

	calParams := make(map[string]float64)
	if service.MaxAcreage != 0 {
		calParams["MaxAcreage"] = service.MaxAcreage
	}
	if service.MinAcreage != 0 {
		calParams["MinAcreage"] = service.MinAcreage
	}
	if service.MaxTotalPrice != 0 {
		calParams["MaxTotalPrice"] = service.MaxTotalPrice
	}
	if service.MinTotalPrice != 0 {
		calParams["MinTotalPrice"] = service.MinTotalPrice
	}
	if service.MaxUnitPrice != 0 {
		calParams["MaxUnitPrice"] = service.MaxUnitPrice
	}
	if service.MinUnitPrice != 0 {
		calParams["MinUnitPrice"] = service.MinUnitPrice
	}
	if service.IsDecoration != 0 {
		calParams["IsDecoration"] = float64(service.IsDecoration)
	}
	if service.IsNotDecoration != 0 {
		calParams["IsNotDecoration"] = float64(service.IsNotDecoration)
	}
	if service.IsNotDecoration != 0 {
		calParams["IsNotDecoration"] = float64(service.IsNotDecoration)
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
	if service.PredictCredDate != 0 {
		calParams["PredictCredDate"] = float64(service.PredictCredDate)
	}
	calParams["needed"] = 1

	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 10
	}
	res := QueryProject(service.Start, size, commonParam, calParams)
	if res != nil {
		var result []model.Project
		for _, item := range res.Each(reflect.TypeOf(model.Project{})) {
			if t, ok := item.(model.Project); ok {
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

//获取楼盘
func (service *GetProjectService) GetProject() serializer.Response {
	res := GetProject(service.ProjectId)
	if res != nil {
		return serializer.Response{
			Code: 200,
			Data: res.Source,
			Msg: "success",
		}
	}
	return serializer.Response{
		Code: 400,
		Data: nil,
		Msg: "未找到数据",
	}
}

//根据楼盘获取开盘
func (service *GetProjectCredService) GetProjectCred() serializer.Response {
	commonParam := make(map[string]string)
	if service.Sort != "" {
		commonParam["sort"] = service.Sort
	} else {
		commonParam["sort"] = "CredDate"
	}
	if service.SortType != "" {
		commonParam["sortType"] = service.SortType
	} else {
		commonParam["sortType"] = "desc"
	}
	commonParam["ProjectId"] = service.ProjectId
	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 10
	}
	res := GetProjectCred(service.Start, size, commonParam)
	if res != nil {
		var result []model.Cred
		for _, item := range res.Each(reflect.TypeOf(model.Cred{})) {
			if t, ok := item.(model.Cred); ok {
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

//根据楼盘获取开盘服务
//func (service *GetCredHouseService) GetCredHouse() serializer.Response {
//	commonParam := make(map[string]string)
//	if service.Sort != "" {
//		commonParam["sort"] = service.Sort
//	} else {
//		commonParam["sort"] = "UpdatedAt"
//	}
//	if service.SortType != "" {
//		commonParam["sortType"] = service.SortType
//	} else {
//		commonParam["sortType"] = "desc"
//	}
//	commonParam["CredId"] = service.CredId
//	var size int = 0
//	if service.Size != 0 {
//		size = service.Size
//	}  else {
//		size = 10
//	}
//	res := GetCredHouse(service.Start, size, commonParam)
//	if res != nil {
//		var result []model.House
//		for _, item := range res.Each(reflect.TypeOf(model.House{})) {
//			if t, ok := item.(model.House); ok {
//				result = append(result, t)
//			}
//		}
//		return serializer.Response{
//			Code: 200,
//			Data: result,
//			Msg: "success",
//		}
//	} else {
//		return serializer.Response{
//			Code: 400,
//			Msg: "暂无数据",
//		}
//	}
//}

//热门搜索
func (service *GetHotService) GetHot() serializer.Response {
	commonParam := make(map[string]string)
	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 5
	}
	commonParam["sort"] = "ViewCount"
	commonParam["sortType"] = "desc"
	calParams := make(map[string]float64)
	calParams["needed"] = 1
	res := QueryProject(service.Start, size, commonParam, calParams)
	if res != nil {
		var result []*HotProject
		for _, item := range res.Each(reflect.TypeOf(model.Project{})) {
			if t, ok := item.(model.Project); ok {
				tempHot := new(HotProject)
				tempHot.ProjectName = t.ProjectName
				tempHot.ProjectId = t.ID
				result = append(result, tempHot)
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

//即将取证时间分组服务
func (service *GetPredictCredDate) GetPredictCredDate() serializer.Response {

	res := repo.NewProjectRepo().GetPredictCredDate()

	if res != nil {
		return serializer.Response{
			Code: 200,
			Data: res,
			Msg: "暂无数据",
		}
	}

	return serializer.Response{
		Code: 200,
		Msg: "暂无数据",
	}
}