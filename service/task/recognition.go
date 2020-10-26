/**
 * @Description:
 * @File: recognition
 * @Date: 2020/8/17 0017 18:39
 */

package task

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"csxft/service/es_update"
)

//正在认筹任务
type RecognitionService struct {
}

//正在认筹到期任务
type NotRecognitionService struct {
}

func (service *RecognitionService) GetRecognitionTask() serializer.Response {

	data, err := repo.NewBatchRepo().GetRecognitionTask()
	if err != nil || len(data) <= 0 {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: "暂无相关数据",
		}
	}
	for i, item := range data {
		//修改db
		model.DB.Model(&data[i]).Update(map[string]interface{}{"status": 3, "is_recognition": 1})
		//修改es
		batchEsParam := make(map[string]interface{})
		batchEsParam["IsRecognition"] = 1
		batchEsParam["Status"] = 3
		batchEsParam["StatusName"] = "正在认筹"
		es_update.Update(&batchEsParam, int(item.ID), "batch")
		projectEsParam := make(map[string]interface{})
		projectEsParam["IsRecognition"] = 1
		//修改楼盘
		project := new(model.Project)
		err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
		if err != nil {
			continue
		}
		model.DB.Model(&project).Update("is_recognition", 1)
		es_update.Update(&projectEsParam, int(item.ProjectId), "project")
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}

func (service *NotRecognitionService) GetNotRecognitionTask() serializer.Response {

	data, err := repo.NewBatchRepo().GetNotRecognitionTask()
	if err != nil || len(data) <= 0 {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: "暂无相关数据",
		}
	}
	for i, item := range data {
		//project 修改参数
		projectEsParam := make(map[string]interface{})
		projectEsParam["IsRecognition"] = 0
		projectDbParams := make(map[string]interface{})
		projectDbParams["IsRecognition"] = 0
		//batch es修改参数
		batchEsParam := make(map[string]interface{})
		batchEsParam["Status"] = 5
		batchEsParam["IsRecognition"] = 0
		//数据库修改参数（批次）
		dbParams := make(map[string]interface{})
		dbParams["status"] = 5
		dbParams["is_recognition"] = 0
		//判断认筹时间过了后有无摇号时间
		if !item.LotteryTime.IsZero() && item.LotteryTime.Unix() > item.SolicitEnd.Unix() {
			dbParams["status"] = 4
			dbParams["is_iottery"] = 1
			projectEsParam["IsIottery"] = 1
			projectDbParams["IsIottery"] = 1
			batchEsParam["IsIottery"] = 1
			batchEsParam["Status"] = 4
		}
		switch dbParams["status"]  {
		case 1:
			batchEsParam["StatusName"] = "即将取证"
			break
		case 2:
			batchEsParam["StatusName"] = "最新取证"
			break
		case 3:
			batchEsParam["StatusName"] = "正在认筹"
			break
		case 4:
			batchEsParam["StatusName"] = "最新摇号"
			break
		case 5:
			batchEsParam["StatusName"] = "在售楼盘"
			break
		}
		//修改批次db
		model.DB.Model(&data[i]).Updates(dbParams)
		//model.DB.Model(&data[i]).Update(map[string]interface{}{"status": 10, "is_recognition": 0})
		//修改楼盘db
		project := new(model.Project)
		err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
		if err != nil {
			continue
		}
		model.DB.Model(&project).Updates(projectDbParams)
		//修改es
		es_update.Update(&batchEsParam, int(item.ID), "batch")
		es_update.Update(&projectEsParam, int(item.ProjectId), "project")
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}


