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
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: err.Error(),
		}
	}
	for _, item := range data {
		batch := new(model.Batch)
		if item.BatchId != 0 {
			err := model.DB.Model(model.Batch{}).Where("id = ?", item.BatchId).First(&batch).Error
			if err != nil {
				continue
			}
		} else {
			continue
		}

		//判断是否为今天取证 是则修改批次状态
		if item.CredDate.Day() == util.GetToday().Day() {
			batch.Status = 2
			model.DB.Model(&batch).Update("status", 2)
		}

		//修改批次 是否为最新取证 用于多状态
		model.DB.Model(&batch).Update("is_will_cred", 1)
		batchEsParam := make(map[string]interface{})
		batchEsParam["IsWillCred"] = 1
		batchEsParam["Status"] = batch.Status
		es_update.Update(&batchEsParam, int(batch.ID), "batch")

		//修改楼盘
		project := new(model.Project)
		err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
		if err != nil {
			continue
		}
		project.IsNewCred = 1
		model.DB.Model(&project).Update("is_will_cred", 1)
		projectEsParam := make(map[string]interface{})
		projectEsParam["IsWillCred"] = 1
		es_update.Update(&projectEsParam, int(project.ID), "project")

	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}


func (service *NotNewCredTaskService) GetNotNewCredTask() serializer.Response {

	data, err := repo.NewCredRepo().GetNotNewCredTask()

	fmt.Println(data)
	fmt.Println(err)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}
