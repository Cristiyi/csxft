/**
 * @Description:
 * @File: iottery_repo
 * @Date: 2020/7/17 0017 14:43
 */

package repo

import "csxft/model"

type LotteryHistoryRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (lotteryHistory *model.LotteryHistory, err error)
}

func NewLotteryHistoryRepo() LotteryHistoryRepo {
	return &lotteryHistoryRepo{}
}

type lotteryHistoryRepo struct {
	thisModel model.LotteryHistory
}

func (c lotteryHistoryRepo) GetToEsData(id uint64) (lotteryHistory *model.LotteryHistory, err error) {
	lotteryHistory = new(model.LotteryHistory)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&lotteryHistory).Error
	return
}

