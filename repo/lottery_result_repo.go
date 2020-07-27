/**
 * @Description:
 * @File: lottery_result
 * @Date: 2020/7/23 0023 14:52
 */

package repo

import (
	"CMD-XuanFangTong-Server/model"
	"CMD-XuanFangTong-Server/util"
)

type LotteryResultRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (lotteryResult *model.LotteryResult, err error)
}

func NewLotteryResultRepo() LotteryResultRepo {
	return &lotteryResultRepo{}
}

type lotteryResultRepo struct {
	thisModel model.LotteryResult
}

func (c lotteryResultRepo) GetToEsData(id uint64) (lotteryResult *model.LotteryResult, err error) {
	lotteryResult = new(model.LotteryResult)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&lotteryResult).Error
	if lotteryResult.Type == 1 {
		lotteryResult.TypeString = "刚需"
	}
	if lotteryResult.Type == 2 {
		lotteryResult.TypeString = "普通"
	}
	if lotteryResult.IdCard != "" {
		lotteryResult.IdCardBack = lotteryResult.IdCard
		lotteryResult.IdCard = util.HideIdCard(lotteryResult.IdCard)
	}
	return
}

