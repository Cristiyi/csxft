/**
 * @Description:
 * @File: lottery
 * @Date: 2020/8/19 0019 11:26
 */

package task

import (
	"csxft/model"
	"csxft/repo"
	"csxft/serializer"
	"csxft/service/es_update"
	"csxft/util"
)

//正在摇号任务
type LotteryService struct {
}

//正在摇号到期
type NotLotteryService struct {
}

func (service *LotteryService) GetLotteryTask() serializer.Response {

	data, err := repo.NewBatchRepo().GetLotteryTask()
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg: "fail",
			Error: err.Error(),
		}
	}
	for i, item := range data {

		dbBatchParams := make(map[string]interface{})
		dbBatchParams["is_iottery"] = 1
		//判断是否为今天取证 是则修改批次状态
		if item.LotteryTime.Day() == util.GetToday().Day() {
			item.Status = 4
			dbBatchParams["status"] = 4
		}

		//修改批次 是否为最新取证 用于多状态
		model.DB.Model(&data[i]).Updates(dbBatchParams)
		batchEsParam := make(map[string]interface{})
		batchEsParam["IsIottery"] = 1
		batchEsParam["Status"] = item.Status
		switch item.Status {
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
		es_update.Update(&batchEsParam, int(item.ID), "batch")

		//修改楼盘
		project := new(model.Project)
		err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
		if err != nil {
			continue
		}

		model.DB.Model(&project).Update("is_iottery", 1)
		projectEsParam := make(map[string]interface{})
		projectEsParam["IsIottery"] = 1
		es_update.Update(&projectEsParam, int(project.ID), "project")

	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}


func (service *LotteryService) GetNotLotteryTask() serializer.Response {

	data, err := repo.NewBatchRepo().GetLotteryTask()
	if err == nil && len(data) > 0 {
		for _, item := range data {
			batch := new(model.Batch)
			//修改批次 是否为最新取证 用于多状态
			model.DB.Model(&batch).Update("is_new_cred", 0)
			batchEsParam := make(map[string]interface{})
			batchEsParam["IsNewCred"] = 0
			es_update.Update(&batchEsParam, int(batch.ID), "batch")

			//修改楼盘
			project := new(model.Project)
			err := model.DB.Model(model.Project{}).Where("id = ?", item.ProjectId).First(&project).Error
			if err != nil {
				continue
			}
			project.IsNewCred = 0
			model.DB.Model(&project).Update("is_new_cred", 0)
			projectEsParam := make(map[string]interface{})
			projectEsParam["IsNewCred"] = 0
			es_update.Update(&projectEsParam, int(project.ID), "project")

		}
	}

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}
}
