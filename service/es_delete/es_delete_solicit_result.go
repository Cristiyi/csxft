/**
 * @Description:
 * @File: es_delete_lottery_result
 * @Date: 2020/8/28 0028 16:22
 */

package es_delete

import "csxft/serializer"

//删除删除认筹结果
type DeleteSolicitResultService struct {
	SolicitResultId  int `form:"solicit_result_id" json:"solicit_result_id" binding:"required"`
}

//删除删除认筹结果
type DeleteBatchSolicitResultService struct {
	BatchId  int `form:"batch_id" json:"batch_id" binding:"required"`
}

//删除认筹结果
func (service *DeleteSolicitResultService) DeleteSolicitResult() serializer.Response {

	DeleteDoc(service.SolicitResultId, "solicit_result")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

//删除认筹结果（批量）
func (service *DeleteBatchSolicitResultService) DeleteBatchSolicitService() serializer.Response {

	DeleteBatchSolicitResult(service.BatchId)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}
