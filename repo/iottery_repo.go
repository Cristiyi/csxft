/**
 * @Description:
 * @File: iottery_repo
 * @Date: 2020/7/17 0017 14:43
 */

package repo

import "CMD-XuanFangTong-Server/model"

type IotteryRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (iottery *model.Iottery, err error)
}

func NewIotteryRepo() IotteryRepo {
	return &iotteryRepo{}
}

type iotteryRepo struct {
	thisModel model.Iottery
}

func (c iotteryRepo) GetToEsData(id uint64) (iottery *model.Iottery, err error) {
	iottery = new(model.Iottery)
	err = model.DB.Model(c.thisModel).Where("id = ?", id).First(&iottery).Error
	return
}

