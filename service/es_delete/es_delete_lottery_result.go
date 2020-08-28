/**
 * @Description:
 * @File: es_delete_lottery_result
 * @Date: 2020/8/28 0028 16:22
 */

package es_delete

import "csxft/serializer"

//删除摇号结果服务
type DeleteLotteryResultService struct {
	LotteryResultId  int `form:"lottery_result_id" json:"lottery_result_id" binding:"required"`
}

//删除摇号结果服务
type DeleteBatchLotteryResultService struct {
	BatchId  int `form:"batch_id" json:"batch_id" binding:"required"`
}

//删除摇号结果
func (service *DeleteLotteryResultService) DeleteLotteryResult() serializer.Response {

	DeleteDoc(service.LotteryResultId, "lottery_result")

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}

//根据开盘删除一房一价
func (service *DeleteBatchLotteryResultService) DeleteBatchLotteryService() serializer.Response {

	DeleteBatchLotteryResult(service.BatchId)

	return serializer.Response{
		Code: 200,
		Msg: "success",
	}

}
