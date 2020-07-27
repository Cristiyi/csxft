/**
 * @Description: house repo
 * @File: house_repo
 * @Date: 2020/7/10 0010 18:24
 */

package repo

import (
	"CMD-XuanFangTong-Server/model"
)

type HouseRepo interface {
	//获取插入到es的数据
	GetToEsData(id uint64) (houses []model.House, err error)
}

func NewHouseRepo() HouseRepo {
	return &houseRepo{}
}

type houseRepo struct {
	thisModel model.House
}

//获取插入到es的数据
func (c houseRepo) GetToEsData(id uint64) (houses []model.House, err error) {
	err = model.DB.Model(c.thisModel).Where("cred_id = ?", id).Find(&houses).Error
	cred := new(model.Cred)
	model.DB.Model(model.Cred{}).Where("id = ?", id).First(&cred)
	for i, item := range houses {
		if item.DecorationId != 0 {
			houses[i].DecorationString = model.GetDecorationNameById(item.DecorationId).Name
		}
		if item.PurposeId != 0 {
			houses[i].PurposeString = model.GetPurposeNameById(item.DecorationId).Name
		}
		houses[i].BuildNo = cred.BuildingNo
	}
	return
}
