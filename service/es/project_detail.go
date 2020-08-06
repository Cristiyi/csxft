/**
 * @Description: 楼盘详情
 * @File: project_detail
 * @Date: 2020/7/15 0015 18:47
 */

package es

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"reflect"
	"strconv"
)

//获取楼盘服务
type ProjectDetailService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	UserId  uint64 `form:"user_id" json:"project_id"`
	Status int32 `form:"status" json:"status"`
}

//获取最新开盘一房一价服务
type NewCredHouseService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	SortType    string `form:"sort_type" json:"sort_type"`
	Sort    string `form:"sort" json:"sort"`
	//FloorNo string `form:"floor_no" json:"floor_no"`
	BuildNo string `form:"build_no" json:"build_no"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
	Status int32 `form:"status" json:"status"`
}

//获取历史摇号服务
type HistoryIotteryService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Start int `form:"start" json:"start"`
	Size int `form:"size" json:"size"`
}

//获取历史摇号服务
type AllBuildNoService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Status int32 `form:"status" json:"status"`
}

//获取时间轴服务
type TimeLineService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Status int32 `form:"status" json:"status"`
}

//时间轴
type TimeLine struct {
	PreSellTime interface{}
	SolicitBegin interface{}
	SolicitEnd interface{}
	SolicitTime interface{}
	LotteryTime interface{}
	LotteryBegin interface{}
	LotteryEnd interface{}
	ChooseHouseBegin interface{}
	ChooseHouseEnd interface{}
}

type TimeLineResult struct {
	TimeLine map[string]interface{}
	Stage int32
}

//楼盘数据检测服务
type DetailCheckService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	Status int32 `form:"status" json:"status"`
}

//楼盘数据检测
type DetailCheckResult struct {
	Solicit bool
	LotteryNotice bool
	LotteryResult bool
	ChooseHouseNotice bool
}

//公告服务
type GetNoticeService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	NoticeType    string `form:"notice_type" json:"notice_type" binding:"required"`
}

//获取楼盘详情
func (service *ProjectDetailService) ProjectDetail() serializer.Response {
	res := GetProject(service.ProjectId)
	if res == nil {
		return serializer.Response{
			Code: 400,
			Data: nil,
			Msg: "未找到数据",
		}
	}
	data := make(map[string]interface{})
	data["detail"] = res.Source
	data["newCred"] = GetTargetBatch(service.ProjectId, service.Status)
	data["follow"] = 0
	data["follow_count"] = 0
	if service.UserId != 0 {
		projectId, _ := strconv.Atoi(service.ProjectId)
		follow, err := repo.NewFollowRepo().Get(service.UserId, uint64(projectId))
		if err == nil {
			data["follow"] = follow.Status
		}

		data["follow_count"] = repo.NewFollowRepo().GetCount(uint64(projectId))
	}
	return serializer.Response{
		Code: 200,
		Data: data,
		Msg: "success",
	}
}

//获取最近开盘的一房一价
func (service NewCredHouseService) GetNewCredHouse() serializer.Response {

	var credIdResult []string
	batch := GetTargetBatch(service.ProjectId, service.Status)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}
	if batch.Creds != nil {
		for _, item := range batch.Creds {
			credIdResult = append(credIdResult, strconv.Itoa(int(item.ID)))
		}
		var houseResult []model.House
		houseParam := make(map[string]string)
		if service.Sort != "" {
			houseParam["sort"] = service.Sort
		} else {
			houseParam["sort"] = "HouseNo"
		}
		if service.SortType != "" {
			houseParam["sortType"] = service.SortType
		} else {
			houseParam["sortType"] = "desc"
		}
		//if service.FloorNo != "" {
		//	houseParam["FloorNo"] = service.FloorNo
		//}
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

//获取历史摇号
func (service HistoryIotteryService) GetHistoryIottery() serializer.Response {
	commonParam := make(map[string]string)
	commonParam["ProjectId"] = service.ProjectId
	commonParam["sort"] = "UpdatedAt"
	commonParam["sortType"] = "desc"
	var size int = 0
	if service.Size != 0 {
		size = service.Size
	}  else {
		size = 10
	}
	res := GetProjectIottery(service.Start, size, commonParam)
	if res != nil {
		var result []model.Iottery
		for _, item := range res.Each(reflect.TypeOf(model.Iottery{})) {
			if t, ok := item.(model.Iottery); ok {
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

//获取所有楼栋
func (service AllBuildNoService) GetAllBuildNo() serializer.Response {

	var result []string
	batch := GetTargetBatch(service.ProjectId, service.Status)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}
	if batch.Creds != nil {
		for _, item := range batch.Creds {
			result = append(result, item.BuildingNo)
		}
		if result != nil {
			return serializer.Response{
				Code: 200,
				Data: result,
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

//获取时间轴
func (service TimeLineService) GetTimeLine() serializer.Response {

	batch := GetTargetBatch(service.ProjectId, service.Status)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: nil,
			Msg: "暂无数据",
		}
	}
	res := make(map[string]interface{})
	if batch.PreSellTime.IsZero() {
		res["PreSellTime"] = nil
	} else {
		res["PreSellTime"] = batch.PreSellTime
	}
	if batch.SolicitBegin.IsZero() {
		res["SolicitBegin"] = nil
	} else {
		res["SolicitBegin"] = batch.SolicitBegin
	}
	if batch.SolicitEnd.IsZero() {
		res["SolicitEnd"] = nil
	} else {
		res["SolicitEnd"] = batch.SolicitEnd
	}
	if batch.SolicitTime.IsZero() {
		res["SolicitTime"] = nil
	} else {
		res["SolicitTime"] = batch.SolicitTime
	}
	if batch.LotteryTime.IsZero() {
		res["LotteryTime"] = nil
	} else {
		res["LotteryTime"] = batch.LotteryTime
	}
	if batch.LotteryBegin.IsZero() {
		res["LotteryBegin"] = nil
	} else {
		res["LotteryBegin"] = batch.LotteryBegin
	}
	if batch.LotteryEnd.IsZero() {
		res["LotteryEnd"] = nil
	} else {
		res["LotteryEnd"] = batch.LotteryEnd
	}
	if batch.ChooseHouseBegin.IsZero() {
		res["ChooseHouseBegin"] = nil
	} else {
		res["ChooseHouseBegin"] = batch.ChooseHouseBegin
	}
	if batch.ChooseHouseEnd.IsZero() {
		res["ChooseHouseEnd"] = nil
	} else {
		res["ChooseHouseEnd"] = batch.ChooseHouseEnd
	}
	stage := calTimeLine(batch)
	timeLineResult := TimeLineResult{
		res,
		stage,
	}
	return serializer.Response{
		Code: 200,
		Data: timeLineResult,
		Msg: "success",
	}
}

//详情检测
func (service DetailCheckService) DetailCheck () serializer.Response {

	batch := GetTargetBatch(service.ProjectId, service.Status)
	data := new(DetailCheckResult)
	if batch == nil {
		return serializer.Response{
			Code: 200,
			Data: data,
			Msg: "暂无数据",
		}
	}
	solicitParam := make(map[string]string)
	solicitParam["ProjectId"] = service.ProjectId
	solicitParam["BatchId"] = strconv.Itoa(int(batch.ID))
	solicitParam["sort"] = "UpdatedAt"
	solicitParam["sortType"] = "asc"
	solicitRes := QuerySolicitResult(1,10, solicitParam)
	if solicitRes != nil && len(solicitRes.Hits.Hits) > 0 {
		data.Solicit = true
	}

	lotteryParam := make(map[string]string)
	lotteryParam["ProjectId"] = service.ProjectId
	lotteryParam["BatchId"] = strconv.Itoa(int(batch.ID))
	lotteryParam["sort"] = "No"
	lotteryParam["sort"] = "UpdatedAt"
	lotteryParam["sortType"] = "asc"
	lotteryRes := QueryLotteryResult(1,10, lotteryParam)
	if lotteryRes != nil && len(lotteryRes.Hits.Hits) > 0 {
		data.LotteryResult = true
	}

	lotteryNoticeParam := make(map[string]string)
	lotteryNoticeParam["ProjectId"] = service.ProjectId
	lotteryNoticeParam["BatchId"] = strconv.Itoa(int(batch.ID))
	lotteryNoticeParam["NoticeType"] = "2"
	lotteryNoticeRes := GetNotice(lotteryNoticeParam)
	if lotteryNoticeRes != nil && len(lotteryNoticeRes.Hits.Hits) > 0 {
		data.LotteryNotice = true
	}

	chooseHouseNoticeParam := make(map[string]string)
	chooseHouseNoticeParam["ProjectId"] = service.ProjectId
	chooseHouseNoticeParam["BatchId"] = strconv.Itoa(int(batch.ID))
	chooseHouseNoticeParam["NoticeType"] = "1"
	chooseHouseNoticeRes := GetNotice(chooseHouseNoticeParam)
	if chooseHouseNoticeRes != nil && len(chooseHouseNoticeRes.Hits.Hits) > 0 {
		data.ChooseHouseNotice = true
	}

	return serializer.Response{
		Code: 200,
		Data: data,
		Msg: "success",
	}
}

//公告
func (service GetNoticeService) GetNotice () serializer.Response {

	noticeParam := make(map[string]string)
	noticeParam["ProjectId"] = service.ProjectId
	noticeParam["NoticeType"] = service.NoticeType
	noticeRes := GetNotice(noticeParam)
	if noticeRes != nil {
		var result []model.Notice
		for _, item := range noticeRes.Each(reflect.TypeOf(model.Notice{})) {
			if t, ok := item.(model.Notice); ok {
				result = append(result, t)
			}
		}
		return serializer.Response{
			Code: 200,
			Data: result[0],
			Msg: "success",
		}
	}

	return serializer.Response{
		Code: 400,
		Msg: "暂无数据",
	}
}