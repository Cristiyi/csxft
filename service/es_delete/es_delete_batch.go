/**
 * @Description:
 * @File: es_delete_batch
 * @Date: 2020/8/18 0018 11:20
 */

package es_delete

import (
	"csxft/model"
	"csxft/serializer"
	"csxft/service/es_update"
)

//删除楼盘服务
type DeleteBatchService struct {
	BatchId  int `form:"batch_id" json:"batch_id" binding:"required"`
	ProjectId  int `form:"project_id" json:"project_id" binding:"required"`
	Status int `form:"status" json:"status"`
}

func (service *DeleteBatchService) DeleteBatch() serializer.Response {

	DeleteDoc(service.BatchId, "batch")
	project := new(model.Project)
	err := model.DB.Model(project).Where("id = ?", service.ProjectId).First(&project).Error
	var col string
	var colEs string
	if err == nil {
		switch service.Status {
		case 1:
			col = "is_will_cred"
			colEs = "IsWillCred"
		case 2:
			col = "is_new_cred"
			colEs = "IsNewCred"
		case 3:
			col = "is_recognition"
			colEs = "IsRecognition"
		case 4:
			col = "is_iottery"
			colEs = "IsIottery"
		case 5:
			col = "is_sell"
			colEs = "IsSell"
		default:
			break
		}
		if col != "" && colEs != "" {
			model.DB.Model(&project).Update(col, 0)
			projectEsParam := make(map[string]interface{})
			projectEsParam[colEs] = 0
			es_update.Update(&projectEsParam, int(project.ID), "project")
		}
	}


	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}


