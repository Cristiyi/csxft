/**
 * @Description: function
 * @File: function
 * @Date: 2020/7/15 0015 18:31
 */

package es

import (
	"csxft/model"
	"csxft/util"
	"reflect"
)

type NewCredResult struct {
	//ProjectName string  //项目名称
	Decoration string  //装修情况
	VerifyMoney string  //验资金额
	SaleScope string  //销售范围
	CredNoScope string  //预售证
	StructureAcreage string  //建筑面积
	AllNo int  //总套数
	ShedNo int //棚改套数
	RigidNo int  //刚需套数
	OrdinaryNo int  //普通套数
	CustomPrice string  //自定义价格
	StatusName string //状态名称
}

//获取楼盘最新开盘批次信息
func GetNewCred(projectId string, status int32) *NewCredResult {

	data := new(NewCredResult)
	commonParam := make(map[string]string)
	commonParam["ProjectId"] = projectId
	commonParam["IsNew"] = "1"
	commonParam["sort"] = "UpdatedAt"
	commonParam["sortType"] = "desc"
	esRes := GetProjectCred(0, 10, commonParam)
	if esRes != nil {
		var result []model.Cred
		for _, item := range esRes.Each(reflect.TypeOf(model.Cred{})) {
			if t, ok := item.(model.Cred); ok {
				result = append(result, t)
			}
		}
		//if result != nil {
		//	maxAcreage := result[0].MaxArea
		//	minAcreage := result[0].MinArea
		//	decoration := result[0].RenovationString
		//	saleScope := result[0].SaleScope
		//	customPrice := result[0].CustomPrice
		//	var allNo int
		//	var shedNo int
		//	var rigidNo int
		//	var OrdinaryNo int
		//	var credNoScope string
		//	var credIdResult string
		//	var allBuildNo string
		//	for i, item := range result {
		//		allNo += item.AllNo
		//		shedNo += item.ShedNo
		//		rigidNo += item.RigidNo
		//		OrdinaryNo += item.OrdinaryNo
		//		if item.MinArea < minAcreage {
		//			minAcreage = item.MinArea
		//		}
		//		if item.MaxArea > maxAcreage {
		//			maxAcreage = item.MaxArea
		//		}
		//		if find := strings.Contains(decoration, item.RenovationString); !find {
		//			decoration = decoration + "," +item.RenovationString
		//		}
		//		if i == len(result)-1 {
		//			credNoScope = credNoScope + item.Cred
		//			saleScope = saleScope + item.BuildingNo
		//			credIdResult = credIdResult + strconv.Itoa(int(item.ID))
		//			allBuildNo = allBuildNo + item.BuildingNo
		//		} else {
		//			credNoScope = credNoScope + item.Cred + ","
		//			saleScope = saleScope + item.BuildingNo + ","
		//			credIdResult = credIdResult + strconv.Itoa(int(item.ID)) + ","
		//			allBuildNo = allBuildNo + item.BuildingNo + ","
		//		}
		//	}
		//	cache.RedisClient.HSet("PROJECT_NEW_CRED", projectId, credIdResult)
		//	cache.RedisClient.HSet("PROJECT_NEW_BUILD_NO", projectId, allBuildNo)
		//	data.Decoration = decoration
		//	data.CredNoScope = credNoScope
		//	data.VerifyMoney = result[0].VerifyMoney
		//	data.SaleScope = saleScope
		//	data.AllNo = allNo
		//	data.RigidNo = rigidNo
		//	data.OrdinaryNo = OrdinaryNo
		//	data.ShedNo = shedNo
		//	data.CustomPrice = customPrice
		//	data.StatusName = result[0].StatusName
		//	if minAcreage == maxAcreage {
		//		data.StructureAcreage = fmt.Sprintf("%.2f",minAcreage)
		//	} else {
		//		data.StructureAcreage = fmt.Sprintf("%.2f",minAcreage) + "-" + fmt.Sprintf("%.2f",maxAcreage)
		//	}
		//}
	}

	return data

}

func calTimeLine(batch *model.Batch) int32 {
	todayTime := util.GetTodayUnix()
	if preSellTimeUnix := batch.PreSellTime.Unix(); preSellTimeUnix != 0{
		if todayTime == preSellTimeUnix {
			return 1
		}
	}
	if solicitBeginUnix := batch.SolicitEnd.Unix(); solicitBeginUnix != 0{
		solicitEndUnix := batch.SolicitEnd.Unix()
		if todayTime>=solicitBeginUnix && todayTime<=solicitEndUnix {
			return 2
		}
		if preSellTimeUnix := batch.PreSellTime.Unix(); preSellTimeUnix != 0{
			if todayTime >= preSellTimeUnix && todayTime < solicitBeginUnix {
				return 1
			}
		}
	}
	if solicitTimeUnix := batch.SolicitTime.Unix(); solicitTimeUnix != 0{
		if todayTime == solicitTimeUnix {
			return 3
		}
	}
	if lotteryTimeUnix := batch.LotteryTime.Unix(); lotteryTimeUnix != 0{
		if todayTime == lotteryTimeUnix {
			return 4
		}
	}
	if lotteryBeginUnix := batch.LotteryBegin.Unix(); lotteryBeginUnix != 0{
		lotteryEndUnix := batch.LotteryEnd.Unix()
		if todayTime >= lotteryBeginUnix && todayTime <= lotteryEndUnix {
			return 5
		}
	}
	if chooseHouseBeginUnix := batch.ChooseHouseBegin.Unix(); chooseHouseBeginUnix != 0{
		chooseHouseEndUnix := batch.ChooseHouseEnd.Unix()
		if todayTime >= chooseHouseBeginUnix && todayTime <= chooseHouseEndUnix {
			return 6
		}
	}
	return 0
}

func GetTargetBatch(projectId string, status int32) *model.Batch {
	batchRes := GetBatch(projectId, status)
	if batchRes != nil && len(batchRes.Hits.Hits) > 0 {
		var batches []model.Batch
		for _, item := range batchRes.Each(reflect.TypeOf(model.Batch{})) {
			if t, ok := item.(model.Batch); ok {
				batches = append(batches, t)
			}
		}
		return &batches[0]
	}
	return nil
}
