/**
 * @Description: 楼盘详情
 * @File: project_detail
 * @Date: 2020/7/15 0015 18:47
 */

package es

import (
	"CMD-XuanFangTong-Server/cache"
	"CMD-XuanFangTong-Server/model"
	"CMD-XuanFangTong-Server/repo"
	"CMD-XuanFangTong-Server/serializer"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

//获取楼盘服务
type ProjectDetailService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
	UserId  uint64 `form:"user_id" json:"project_id" binding:"required"`
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
}

//获取时间轴服务
type TimeLineService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
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
	TimeLine TimeLine
	Stage int32
}

//楼盘数据检测服务
type DetailCheckService struct {
	ProjectId    string `form:"project_id" json:"project_id" binding:"required"`
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

//获取楼盘
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

	newCred := GetNewCred(service.ProjectId)

	data["follow"] = 0
	projectId, _ := strconv.Atoi(service.ProjectId)
	follow, err := repo.NewFollowRepo().Get(service.UserId, uint64(projectId))
	if err == nil {
		data["follow"] = follow.Status
	}

	data["follow_count"] = repo.NewFollowRepo().GetCount(uint64(projectId))

	if newCred != nil {
		data["newCred"] = newCred
	}

	return serializer.Response{
		Code: 200,
		Data: data,
		Msg: "success",
	}
}

//获取最近开盘的一房一价
func (service NewCredHouseService) GetNewCredHouse() serializer.Response {

	//获取最新的预售信息 先从redis取，取不到去es查询
	var credIdResult []string
	rs, err := cache.RedisClient.HGet("PROJECT_NEW_CRED", service.ProjectId).Result()
	if err != nil {
		credParam := make(map[string]string)
		credParam["ProjectId"] = service.ProjectId
		credParam["IsNew"] = "1"
		credParam["sort"] = "UpdatedAt"
		credParam["sortType"] = "desc"
		credRes := GetProjectCred(0, 10, credParam)
		if credRes != nil {
			for _, item := range credRes.Each(reflect.TypeOf(model.Cred{})) {
				if t, ok := item.(model.Cred); ok {
					credIdResult = append(credIdResult, strconv.Itoa(int(t.ID)))
				}
			}
		}
	} else {
		credIdResult = strings.Split(rs, ",")
	}
	if credIdResult != nil {
		//根据最新的预售信息的id范围获取房屋信息
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
		Code: 400,
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

	rs, err := cache.RedisClient.HGet("PROJECT_NEW_BUILD_NO", service.ProjectId).Result()
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg: "暂无数据",
		}
	} else {
		result := strings.Split(rs, ",")
		return serializer.Response{
			Code: 200,
			Data: result,
			Msg: "success",
		}
	}

}

//获取时间轴
func (service TimeLineService) GetTimeLine() serializer.Response {

	//获取最新的预售信息 先从redis取，取不到去es查询
	var credIdResult []string
	rs, err := cache.RedisClient.HGet("PROJECT_NEW_CRED", service.ProjectId).Result()
	if err != nil {
		credParam := make(map[string]string)
		credParam["ProjectId"] = service.ProjectId
		credParam["IsNew"] = "1"
		credParam["sort"] = "UpdatedAt"
		credParam["sortType"] = "desc"
		credRes := GetProjectCred(0, 10, credParam)
		if credRes != nil {
			for _, item := range credRes.Each(reflect.TypeOf(model.Cred{})) {
				if t, ok := item.(model.Cred); ok {
					credIdResult = append(credIdResult, strconv.Itoa(int(t.ID)))
				}
			}
		}
	} else {
		credIdResult = strings.Split(rs, ",")
	}

	if credIdResult != nil {
		credRes := GetCred(credIdResult[0]).Source
		var res map[string]interface{}
		cred := new(model.Cred)
		json.Unmarshal([]byte(credRes), &res)
		json.Unmarshal([]byte(credRes), &cred)
		if cred.PreSellTime.IsZero() {
			res["PreSellTime"] = nil
		}
		if cred.SolicitBegin.IsZero() {
			res["SolicitBegin"] = nil
		}
		if cred.SolicitEnd.IsZero() {
			res["SolicitEnd"] = nil
		}
		if cred.SolicitTime.IsZero() {
			res["SolicitTime"] = nil
		}
		if cred.LotteryTime.IsZero() {
			res["LotteryTime"] = nil
		}
		if cred.LotteryBegin.IsZero() {
			res["LotteryBegin"] = nil
		}
		if cred.LotteryEnd.IsZero() {
			res["LotteryEnd"] = nil
		}
		if cred.ChooseHouseBegin.IsZero() {
			res["ChooseHouseBegin"] = nil
		}
		if cred.ChooseHouseEnd.IsZero() {
			res["ChooseHouseEnd"] = nil
		}
		if res != nil {
			timeLine := TimeLine{
				PreSellTime: res["PreSellTime"],
				SolicitBegin: res["SolicitBegin"],
				SolicitEnd: res["SolicitEnd"],
				SolicitTime: res["SolicitTime"],
				LotteryTime: res["LotteryTime"],
				LotteryBegin: res["LotteryBegin"],
				LotteryEnd: res["LotteryEnd"],
				ChooseHouseBegin: res["ChooseHouseBegin"],
				ChooseHouseEnd: res["ChooseHouseEnd"],
			}
			stage := calTimeLine(cred)
			timeLineResult := TimeLineResult{
				timeLine,
				stage,
			}
			return serializer.Response{
				Code: 200,
				Data: timeLineResult,
				Msg: "success",
			}
		}
	}

	timeLine := TimeLine{
		PreSellTime: nil,
		SolicitBegin: nil,
		SolicitEnd: nil,
		SolicitTime: nil,
		LotteryTime: nil,
		LotteryBegin: nil,
		LotteryEnd: nil,
		ChooseHouseBegin: nil,
		ChooseHouseEnd: nil,
	}

	return serializer.Response{
		Code: 200,
		Data: timeLine,
		Msg: "success",
	}
}

//详情检测
func (service DetailCheckService) DetailCheck () serializer.Response {

	data := new(DetailCheckResult)
	solicitParam := make(map[string]string)
	solicitParam["ProjectId"] = service.ProjectId
	solicitParam["IsNew"] = "1"
	solicitParam["sort"] = "UpdatedAt"
	solicitParam["sortType"] = "asc"
	solicitRes := QuerySolicitResult(1,10, solicitParam)
	if solicitRes != nil && len(solicitRes.Hits.Hits) > 0 {
		data.Solicit = true
	}

	lotteryParam := make(map[string]string)
	lotteryParam["ProjectId"] = service.ProjectId
	lotteryParam["IsNew"] = "1"
	lotteryParam["sort"] = "No"
	lotteryParam["sort"] = "UpdatedAt"
	lotteryParam["sortType"] = "asc"
	lotteryRes := QueryLotteryResult(1,10, solicitParam)
	if lotteryRes != nil && len(solicitRes.Hits.Hits) > 0 {
		data.LotteryResult = true
	}

	lotteryNoticeParam := make(map[string]string)
	lotteryNoticeParam["ProjectId"] = service.ProjectId
	lotteryNoticeParam["NoticeType"] = "1"
	lotteryNoticeRes := GetNotice(lotteryNoticeParam)
	if lotteryNoticeRes != nil && len(solicitRes.Hits.Hits) > 0 {
		data.LotteryNotice = true
	}

	chooseHouseNoticeParam := make(map[string]string)
	chooseHouseNoticeParam["ProjectId"] = service.ProjectId
	chooseHouseNoticeParam["NoticeType"] = "2"
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