/**
 * @Description:
 * @File: lottery_result
 * @Date: 2020/7/23 0023 14:52
 */

package repo

import (
	"csxft/model"
	"csxft/util"
)

type LotteryResultRepo interface {
	//获取插入到es的数据
	GetToEsData(batchId uint64) (lotteryResult []*model.LotteryResult, err error)
}

func NewLotteryResultRepo() LotteryResultRepo {
	return &lotteryResultRepo{}
}

type lotteryResultRepo struct {
	thisModel model.LotteryResult
}

func (c lotteryResultRepo) GetToEsData(id uint64) (lotteryResult []*model.LotteryResult, err error) {

	err = model.DB.Model(c.thisModel).Where("batch_id = ?", id).First(&lotteryResult).Error
	for i, item := range lotteryResult {
		if item.Type == 1 {
			lotteryResult[i].TypeString = "刚需"
		}
		if item.Type == 2 {
			lotteryResult[i].TypeString = "普通"
		}
		if item.IdCard != "" {
			lotteryResult[i].IdCardBack = lotteryResult[i].IdCard
			lotteryResult[i].IdCard = util.HideIdCard(lotteryResult[i].IdCard)
		}
	}
	return
}

