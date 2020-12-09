/**
 * @Description:
 * @File: new_cred
 * @Date: 2020/8/17 0017 14:45
 */

package task

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"csxft/service/es_update"
	"csxft/util"
	"fmt"
)

//最新取证任务
type NewCredTaskService struct {
}

//取消最新取证任务
type NotNewCredTaskService struct {
}

func (service *NewCredTaskService) GetNewCredTask() serializer.Response {

	data, err := repo.NewCredRepo().GetNewCredTask()
	if err != nil || len(data) <= 0 {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: "暂无相关数据",
		}
	}
	for _, item := range data {
		batch := new(model.Batch)
		if item.BatchId != 0 {
			err := model.DB.Model(model.Batch{}).Where("id = ?", item.BatchId).First(&batch).Error
			if err != nil  || batch.Status != 1{
				continue
			}
		} else {
			continue
		}
		dbBatchParams := make(map[string]interface{})
		dbBatchParams["is_new_cred"] = 1
		dbBatchParams["is_will_cred"] = 0
		//判断是否为今天取证 是则修改批次状态
		if item.CredDate.Day() == util.GetToday().Day() {
			dbBatchParams["status"] = 2
		}

		//修改批次 是否为最新取证 用于多状态
		model.DB.Model(&batch).Updates(dbBatchParams)
		batchEsParam := make(map[string]interface{})
		batchEsParam["IsNewCred"] = 1
		batchEsParam["Status"] = 2
		batchEsParam["IsWillCred"] = 0
		switch batch.Status {
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
		es_update.Update(&batchEsParam, int(batch.ID), "batch")

		//修改楼盘
		project := new(model.Project)
		err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
		if err != nil {
			continue
		}
		dbProjectParams := make(map[string]interface{})
		dbProjectParams["is_new_cred"] = 1
		dbProjectParams["is_will_cred"] = 0
		model.DB.Model(&project).Updates(dbProjectParams)
		projectEsParam := make(map[string]interface{})
		projectEsParam["IsNewCred"] = 1
		projectEsParam["IsWillCred"] = 0
		es_update.Update(&projectEsParam, int(project.ID), "project")

	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}


func (service *NotNewCredTaskService) GetNotNewCredTask() serializer.Response {

	data, err := repo.NewCredRepo().GetNotNewCredTask()
	if err != nil || len(data) <= 0 {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: "暂无相关数据",
		}
	}
	for _, item := range data {
		batch := new(model.Batch)
		if item.BatchId != 0 {
			err := model.DB.Model(model.Batch{}).Where("id = ?", item.BatchId).First(&batch).Error
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			continue
		}
		//修改批次 是否为最新取证 用于多状态
		dbBatchParams := make(map[string]interface{})
		dbBatchParams["is_new_cred"] = 0
		dbBatchParams["status"] = 5
		model.DB.Model(&batch).Updates(dbBatchParams)

		batchEsParam := make(map[string]interface{})
		batchEsParam["IsNewCred"] = 0
		batchEsParam["Status"] = 5
		batchEsParam["StatusName"] = "在售楼盘"
		es_update.Update(&batchEsParam, int(batch.ID), "batch")

		//修改楼盘
		project := new(model.Project)
		err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
		if err != nil {
			continue
		}
		//project.IsNewCred = 0
		model.DB.Model(&project).Update("is_new_cred", 0)
		projectEsParam := make(map[string]interface{})
		projectEsParam["IsNewCred"] = 0
		es_update.Update(&projectEsParam, int(project.ID), "project")

	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}
